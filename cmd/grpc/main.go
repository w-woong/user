package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"runtime"
	"syscall"
	"time"

	"github.com/go-wonk/si"
	"github.com/go-wonk/si/sigorm"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/w-woong/common"
	commonadapter "github.com/w-woong/common/adapter"
	"github.com/w-woong/common/configs"
	userpb "github.com/w-woong/common/dto/protos/user/v1"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/common/middlewares"
	commonport "github.com/w-woong/common/port"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/common/wrapper"
	"github.com/w-woong/user/adapter"
	"github.com/w-woong/user/delivery"
	"github.com/w-woong/user/entity"
	"github.com/w-woong/user/port"
	"github.com/w-woong/user/usecase"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var (
	Version = "undefined"

	printVersion     bool
	tickIntervalSec  int = 30
	addr             string
	certPem, certKey string
	readTimeout      int
	writeTimeout     int
	configName       string
	maxProc          int

	usePprof  = false
	pprofAddr = ":56060"
)

func init() {
	flag.StringVar(&addr, "addr", ":42001", "listen address")
	flag.BoolVar(&printVersion, "version", false, "print version")
	flag.IntVar(&tickIntervalSec, "tick", 30, "tick interval in second")
	flag.StringVar(&certKey, "key", "./certs/server.key", "server key")
	flag.StringVar(&certPem, "pem", "./certs/server.crt", "server pem")
	flag.IntVar(&readTimeout, "readTimeout", 30, "read timeout")
	flag.IntVar(&writeTimeout, "writeTimeout", 30, "write timeout")
	flag.StringVar(&configName, "config", "./configs/server.yml", "config file name")
	flag.IntVar(&maxProc, "mp", runtime.NumCPU(), "GOMAXPROCS")

	flag.BoolVar(&usePprof, "pprof", false, "use pprof")
	flag.StringVar(&pprofAddr, "pprof_addr", ":56060", "pprof listen address")

	flag.Parse()
}

func main() {
	var err error

	if printVersion {
		fmt.Printf("version \"%v\"\n", Version)
		return
	}
	runtime.GOMAXPROCS(maxProc)

	// config
	conf := common.Config{}
	if err := configs.ReadConfigInto(configName, &conf); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// logger
	logger.Open(conf.Logger.Level, conf.Logger.Stdout,
		conf.Logger.File.Name, conf.Logger.File.MaxSize, conf.Logger.File.MaxBackup,
		conf.Logger.File.MaxAge, conf.Logger.File.Compressed)
	defer logger.Close()

	// db
	sqlDB, err := si.OpenSqlDB(conf.Server.Repo.Driver, conf.Server.Repo.ConnStr,
		conf.Server.Repo.MaxIdleConns, conf.Server.Repo.MaxOpenConns,
		time.Duration(conf.Server.Repo.ConnMaxLifetimeMinutes)*time.Minute)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer sqlDB.Close()

	// gorm
	var gormDB *gorm.DB
	switch conf.Server.Repo.Driver {
	case "pgx":
		gormDB, err = sigorm.OpenPostgres(sqlDB)
	default:
		logger.Error(conf.Server.Repo.Driver + " is not allowed")
		os.Exit(1)
	}
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	gormDB.AutoMigrate(&entity.User{}, &entity.Email{}, &entity.Password{}, &entity.Personal{})

	var txBeginner common.TxBeginner
	var pwRepo port.PasswordRepo
	var userRepo port.UserRepo
	switch conf.Server.Repo.Driver {
	case "pgx":
		txBeginner = txcom.NewGormTxBeginner(gormDB)
		userRepo = adapter.NewPgUser(gormDB)
		pwRepo = adapter.NewPasswordPg(gormDB)

	default:
		logger.Error(conf.Server.Repo.Driver + " is not allowed")
		os.Exit(1)
	}
	userUsc := usecase.NewUser(txBeginner, userRepo, pwRepo)

	idTokenValidators := make(commonport.IDTokenValidators)
	for _, v := range conf.Client.Oauth2.IDTokenValidators {
		if v.Type == "jwks" {
			jwksUrl, err := commonadapter.GetJwksUrl(v.OpenIDConfUrl)
			if err != nil {
				logger.Error(err.Error())
				os.Exit(1)
			}
			validator := commonadapter.NewJwksIDTokenValidator(jwksUrl, v.Token.TokenSourceKeyName, v.Token.IDKeyName, v.Token.IDTokenKeyName)
			idTokenValidators[v.Token.Source] = validator
		}
	}

	// grpc
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	svr, err := wrapper.NewGrpcServer(conf.Server.Grpc, certPem, certKey, false,
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middlewares.AuthIDTokenInterceptor(idTokenValidators),
		)))
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// grpc server
	userGrpcServer := delivery.NewUserGrpcServer(time.Duration(conf.Server.Grpc.Timeout)*time.Second, userUsc)
	userpb.RegisterUserServiceServer(svr, userGrpcServer)

	grpcServer := wrapper.GrpcServer{Svr: svr}

	// ticker
	ticker := time.NewTicker(time.Duration(tickIntervalSec) * time.Second)
	tickerDone := make(chan bool)
	common.StartTicker(tickerDone, ticker, func(t time.Time) {
		logger.Info(fmt.Sprintf("NoOfGR:%v, %v", runtime.NumGoroutine(), t))
	})

	// signal, wait for it to shutdown http server.
	common.StartSignalStopper(&grpcServer, syscall.SIGINT, syscall.SIGTERM)

	// start
	logger.Info("start listening on " + addr)
	if err = grpcServer.Serve(lis); err != nil {
		logger.Error(err.Error())
	}

	// finish
	ticker.Stop()
	tickerDone <- true
	logger.Info("finished")
}
