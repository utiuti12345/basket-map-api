package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func SetupDB() *gorm.DB {
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	// 接続
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Tokyo", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// デバッグの設定
	//debugMode, err := strconv.ParseBool(os.Getenv("DB_DEBUG_MODE"))
	//if err != nil {
	//	panic(err.Error())
	//}

	return db
}
