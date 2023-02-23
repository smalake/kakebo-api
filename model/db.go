package model

import (
	"os"

	"github.com/smalake/kakebo-api/utils/logging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB接続
// func ConnectDB() *sql.DB {
// 	jst, err := time.LoadLocation("Asia/Tokyo")
// 	if err != nil {
// 		logging.WriteErrorLog(err.Error(), true)
// 	}

// 	c := mysql.Config{
// 		DBName:               os.Getenv("DB_NAME"),
// 		User:                 os.Getenv("DB_USER"),
// 		Passwd:               os.Getenv("DB_PASSWORD"),
// 		Addr:                 os.Getenv("DB_ADDRESS"),
// 		Net:                  "tcp",
// 		ParseTime:            true,
// 		AllowNativePasswords: true,
// 		Collation:            "utf8mb4_unicode_ci",
// 		Loc:                  jst,
// 	}
// 	db, err := sql.Open("mysql", c.FormatDSN())

// 	if err != nil {
// 		logging.WriteErrorLog(err.Error(), true)
// 	}

//		return db
//	}
func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		panic("failed to connect database")
	}
	return db
}
