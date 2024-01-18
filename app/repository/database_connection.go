package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/riny/demo-go-gin/app/model"
	"log/slog"
	"os"
)

func NewDb() (db *sql.DB, err error) {
	dbConn := model.DataBase{Dsn: os.Getenv("DSN")}

	dbConn.Db, err = sql.Open("mysql", dbConn.Dsn)
	if err != nil {
		return nil, err
	}

	return dbConn.Db, nil
}

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
}
