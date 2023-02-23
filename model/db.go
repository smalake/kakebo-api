package model

import (
	"database/sql"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/smalake/kakebo-api/utils/logging"
)

// DB接続
func ConnectDB() *sql.DB {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
	}

	c := mysql.Config{
		DBName:               os.Getenv("DB_NAME"),
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Addr:                 os.Getenv("DB_ADDRESS"),
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  jst,
	}
	db, err := sql.Open("mysql", c.FormatDSN())

	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
	}

	return db
}
