package benchmark

import (
	"log"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   uint
	Name string
}

func insertRecord(b *testing.B, db *gorm.DB) {
	user := User{Name: "Test User"}
	if err := db.Create(&user).Error; err != nil {
		b.Fatal(err)
	}
}

// func BenchmarkMaxOpenConns1(b *testing.B) {
// 	dsn := "host=localhost user=postgres password=postgres dbname=shopdevgo port=5432 sslmode=disable"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("failed to connect to database: %v", err)
// 	}
// 	if db.Migrator().HasTable(&User{}) {
// 		if err := db.Migrator().DropTable(&User{}); err != nil {
// 			b.Fatalf("failed to drop table: %v", err)
// 		}
// 	}
// 	db.AutoMigrate(&User{})
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		b.Fatalf("failed to get sql.DB: %v", err)
// 	}
// 	sqlDB.SetMaxOpenConns(1)
// 	defer sqlDB.Close()

// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			insertRecord(b, db)
// 		}
// 	})
// }

func BenchmarkMaxOpenConns10(b *testing.B) {
	dsn := "host=localhost user=postgres password=postgres dbname=shopdevgo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	if db.Migrator().HasTable(&User{}) {
		if err := db.Migrator().DropTable(&User{}); err != nil {
			b.Fatalf("failed to drop table: %v", err)
		}
	}
	db.AutoMigrate(&User{})
	sqlDB, err := db.DB()
	if err != nil {
		b.Fatalf("failed to get sql.DB: %v", err)
	}
	sqlDB.SetMaxOpenConns(10)
	defer sqlDB.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}
