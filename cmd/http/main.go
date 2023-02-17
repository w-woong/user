package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/go-wonk/si"
	"github.com/go-wonk/si/sigorm"
	"github.com/go-wonk/si/sihttp"
	"github.com/gorilla/mux"
	"github.com/w-woong/common"
	commonadapter "github.com/w-woong/common/adapter"
	"github.com/w-woong/common/configs"
	"github.com/w-woong/common/logger"
	commonport "github.com/w-woong/common/port"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/common/utils"
	"github.com/w-woong/user/adapter"
	"github.com/w-woong/user/cmd/route"
	"github.com/w-woong/user/entity"
	"github.com/w-woong/user/port"
	"github.com/w-woong/user/usecase"
	"gorm.io/gorm"

	// "go.elastic.co/apm/module/apmgorilla/v2"
	postgresapm "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres" // postgres with gorm
	// _ "go.elastic.co/apm/module/apmsql/v2/pq" // postgres sql with pq
	"go.elastic.co/apm/v2"
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

	usePprof    = false
	pprofAddr   = ":56060"
	autoMigrate = false
)

func init() {
	flag.StringVar(&addr, "addr", ":49007", "listen address")
	flag.BoolVar(&printVersion, "version", false, "print version")
	flag.IntVar(&tickIntervalSec, "tick", 60, "tick interval in second")
	flag.StringVar(&certKey, "key", "./certs/key.pem", "server key")
	flag.StringVar(&certPem, "pem", "./certs/cert.pem", "server pem")
	flag.IntVar(&readTimeout, "readTimeout", 30, "read timeout")
	flag.IntVar(&writeTimeout, "writeTimeout", 30, "write timeout")
	flag.StringVar(&configName, "config", "./configs/server.yml", "config file name")
	flag.IntVar(&maxProc, "mp", runtime.NumCPU(), "GOMAXPROCS")

	flag.BoolVar(&usePprof, "pprof", false, "use pprof")
	flag.StringVar(&pprofAddr, "pprof_addr", ":56060", "pprof listen address")
	flag.BoolVar(&autoMigrate, "autoMigrate", false, "auto migrate")

	flag.Parse()
}

func main() {
	var err error

	if printVersion {
		fmt.Printf("version \"%v\"\n", Version)
		return
	}
	runtime.GOMAXPROCS(maxProc)

	// apm
	apmActive, _ := strconv.ParseBool(os.Getenv("ELASTIC_APM_ACTIVE"))
	if apmActive {
		tracer := apm.DefaultTracer()
		defer tracer.Flush(nil)
	}

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

	// gorm
	var gormDB *gorm.DB
	switch conf.Server.Repo.Driver {
	case "pgx":
		// // db
		// sqlDB, err := si.OpenSqlDB(conf.Server.Repo.Driver, conf.Server.Repo.ConnStr,
		// 	conf.Server.Repo.MaxIdleConns, conf.Server.Repo.MaxOpenConns,
		// 	time.Duration(conf.Server.Repo.ConnMaxLifetimeMinutes)*time.Minute)
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }
		// defer sqlDB.Close()
		// gormDB, err = sigorm.OpenPostgres(sqlDB)
		// if err != nil {
		// 	logger.Error(err.Error())
		// 	os.Exit(1)
		// }

		if apmActive {
			gormDB, err = gorm.Open(postgresapm.Open(conf.Server.Repo.ConnStr),
				&gorm.Config{Logger: logger.OpenGormLogger(conf.Server.Repo.LogLevel)},
			)
			if err != nil {
				logger.Error(err.Error())
				os.Exit(1)
			}
			db, err := gormDB.DB()
			if err != nil {
				logger.Error(err.Error())
				os.Exit(1)
			}
			defer db.Close()
		} else {
			// db
			// var db *sql.DB
			db, err := si.OpenSqlDB(conf.Server.Repo.Driver, conf.Server.Repo.ConnStr,
				conf.Server.Repo.MaxIdleConns, conf.Server.Repo.MaxOpenConns, time.Duration(conf.Server.Repo.ConnMaxLifetimeMinutes)*time.Minute)
			if err != nil {
				logger.Error(err.Error())
				os.Exit(1)
			}
			defer db.Close()

			gormDB, err = sigorm.OpenPostgresWithConfig(db,
				&gorm.Config{Logger: logger.OpenGormLogger(conf.Server.Repo.LogLevel)},
			)
			if err != nil {
				logger.Error(err.Error())
				os.Exit(1)
			}
		}
	default:
		logger.Error(conf.Server.Repo.Driver + " is not allowed")
		os.Exit(1)
	}

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

	if autoMigrate {
		gormDB.AutoMigrate(&entity.User{},
			&entity.Email{}, &entity.CredentialPassword{}, &entity.CredentialToken{},
			&entity.Personal{},
			&entity.DeliveryAddress{}, &entity.DeliveryRequest{}, &entity.DeliveryRequestType{},
			&entity.PaymentType{}, &entity.PaymentMethod{})
	}

	userUsc := usecase.NewUser(txBeginner, userRepo, pwRepo)

	var idTokenParser commonport.IDTokenParser
	for _, v := range conf.Client.OAuth2 {
		jwksUrl, err := utils.GetJwksUrl(v.OpenIDConfUrl)
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}

		jwksStore, err := utils.NewJwksCache(jwksUrl)
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
		idTokenParser = commonadapter.NewJwksIDTokenParser(jwksStore)
	}
	tokenCookie := commonadapter.NewTokenCookie(1*time.Hour, conf.Client.IDTokenCookie)

	// http handler
	// userHandler = delivery.NewUserHttpHandler(defaultTimeout, userUsc)
	router := mux.NewRouter()
	// SetRoute(router, conf.Server.Http, idTokenValidators)
	route.UserRoute(router, conf.Server.Http, tokenCookie, idTokenParser, userUsc)

	// http server
	tlsConfig := sihttp.CreateTLSConfigMinTls(tls.VersionTLS12)
	httpServer := sihttp.NewServerCors(router, tlsConfig, addr,
		time.Duration(writeTimeout)*time.Second, time.Duration(readTimeout)*time.Second,
		certPem, certKey,
		strings.Split(conf.Server.Http.AllowedOrigins, ","),
		strings.Split(conf.Server.Http.AllowedHeaders, ","),
		strings.Split(conf.Server.Http.AllowedMethods, ","),
	)

	// ticker
	ticker := time.NewTicker(time.Duration(tickIntervalSec) * time.Second)
	tickerDone := make(chan bool)
	common.StartTicker(tickerDone, ticker, func(t time.Time) {
		logger.Info(fmt.Sprintf("NoOfGR:%v, %v", runtime.NumGoroutine(), t))
	})

	// signal, wait for it to shutdown http server.
	common.StartSignalStopper(httpServer, syscall.SIGINT, syscall.SIGTERM)

	// start
	logger.Info("start listening on " + addr)
	if err = httpServer.Start(); err != nil {
		logger.Error(err.Error())
	}

	// finish
	ticker.Stop()
	tickerDone <- true
	logger.Info("finished")
}
