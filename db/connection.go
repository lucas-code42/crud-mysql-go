package db

import (
	"database/sql"
	"os"

	// Import o Driver do MySql
	// Por padrão o Go vai reclamar do pacote pois possui oiutras dependencias que não estarão explicitamente aqui
	// Então ele pede um comentário de justificativa e tb um "apelido"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectionMySql() (*sql.DB, error) {
	user := os.Getenv("USER")
	pass := os.Getenv("PASS")
	dataBase := os.Getenv("DB")
	urlConnection := user + ":" + pass + "@/" + dataBase + "?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", urlConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	
	return db, nil
}
