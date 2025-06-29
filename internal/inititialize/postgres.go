package inititialize

import (
	"fmt"
	"time"

	"github.com/quockhanhcao/ecommerce/global"
	"github.com/quockhanhcao/ecommerce/internal/entity"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() {
	config := global.Config.PostgresConfig
	dsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=%s"
	connectionString := fmt.Sprintf(dsn, config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode)
	db := openDB(connectionString)
	global.Logger.Info("Connected to Postgres")
	global.DB = db

	// set pool
	SetPool()
	migrateTables()
}

func openDB(connString string) *gorm.DB {
	count := 0
	for {
		db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
		if err != nil {
			global.Logger.Error("Failed to connect to Postgres, retrying...", zap.Error(err))
			count++
		} else {
			global.Logger.Info("Successfully connected to Postgres")
			return db
		}
		if count > 10 {
			global.Logger.Error("Failed to connect to Postgres after 10 attempts, exiting")
			panic(fmt.Errorf("%s: %w", "failed to connect to postgres", err))
		}
		global.Logger.Info("Retrying connection to Postgres in 2 seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}

// func checkErrorPanic(err error, errString string) {
// 	if err != nil {
// 		global.Logger.Error(errString, zap.Error(err))
// 		panic(fmt.Errorf("%s: %w", errString, err))
// 	}
// }

func SetPool() {
	config := global.Config.PostgresConfig
	sqlDb, err := global.DB.DB()
	if err != nil {
		fmt.Println("postgres set pool error:", err)
	}
	sqlDb.SetMaxIdleConns(config.MaxIdleConns)
	sqlDb.SetMaxOpenConns(config.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Second)
}

func migrateTables() {
	err := global.DB.AutoMigrate(
		&entity.User{},
		&entity.Role{},
	)
	if err != nil {
		global.Logger.Error("Failed to migrate tables", zap.Error(err))
	}
}
