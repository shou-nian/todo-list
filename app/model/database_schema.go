package model

import "database/sql"

type DataBase struct {
	Dsn string
	Db  *sql.DB
}
