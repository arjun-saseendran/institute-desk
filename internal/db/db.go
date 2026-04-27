package db

import (
	"log"

	"github.com/arjun-saseendran/institute/internal/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	dsn := "host=localhost user=postgres password=postgres dbname=institute port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {

		return nil, err
	}

	log.Println("Successfully connected to Postgres!")

	err = db.AutoMigrate(&user.User{})
	if err != nil {
		return nil, err
	}

	log.Println("Database migration completed successfully!")

	return db, nil
}
