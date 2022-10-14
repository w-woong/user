package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/w-woong/user/adapter"
	"github.com/w-woong/user/delivery"
	"github.com/w-woong/user/dto"
	"github.com/w-woong/user/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":8080", "listen address")
	flag.Parse()
}

func main() {
	defaultTimeout := 6 * time.Second
	dsn := "host=testpghost user=test password=test123 dbname=woong_user port=5432 sslmode=disable TimeZone=Asia/Seoul"
	var sqlDB *sql.DB
	var err error
	if sqlDB, err = sql.Open("pgx", dsn); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer sqlDB.Close()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(15 * time.Minute)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	db.AutoMigrate(&dto.User{}, &dto.UserEmail{})

	// userRepo := adapter.NewNopUser()
	txBeginner := adapter.NewGormTxBeginner(db)
	userRepo := adapter.NewPgUser(db)
	userUsc := usecase.NewUser(txBeginner, userRepo, defaultTimeout)
	userDelivery := delivery.NewUserHttpHandler(userUsc)

	router := mux.NewRouter()
	router.HandleFunc("/v1/user/register", userDelivery.HandleRegisterUser).Methods("POST")
	router.HandleFunc("/v1/user/{id}", userDelivery.HandleFindByID).Methods("GET")
	router.HandleFunc("/v1/user/{id}", userDelivery.HandleModifyUser).Methods("PUT")
	router.HandleFunc("/v1/user/{id}", userDelivery.HandleRemoveUser).Methods("DELETE")

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	go startSyscallChecker(&server)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}

func startSyscallChecker(hs *http.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigs
	switch sig {
	case syscall.SIGINT:
		log.Println("syscall.SIGINT")
	case syscall.SIGTERM:
		log.Println("syscall.SIGTERM")
	default:
		log.Printf("signal %v", sig)
	}
	hs.Shutdown(context.Background())

}
