package adapter_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/tj/assert"
	"github.com/w-woong/user/adapter"
	"github.com/w-woong/user/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
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

	txBeginner := adapter.NewGormTxBeginner(db)
	userRepo := adapter.NewPgUser(db)

	tx, err := txBeginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()
	_, err = userRepo.CreateUser(context.Background(), tx, entity.User{
		ID:        uuid.New().String(),
		LoginID:   "wonk",
		FirstName: "wonk",
		LastName:  "sun",
		BirthDate: time.Now(),
	})
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())
}

func TestUpdateUser(t *testing.T) {
	dsn := "host=testpghost user=test password=test123 dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Seoul"
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

	userRepo := adapter.NewPgUser(db)

	_, err = userRepo.UpdateUserByID("85bf6aeb-459c-445a-be1e-0b67b8c100ef", entity.User{
		ID:      "85bf6aeb-459c-445a-be1e-0b67b8c100ef",
		LoginID: "wonksing",
	})
	assert.Nil(t, err)
}

func TestBirthDate(t *testing.T) {
	fmt.Println(time.Now().Round(0))
	fmt.Println(time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local))
}
