package db

import (
	"crud-api-mysql/config"
	"database/sql"

	// Import o Driver do MySql
	// Por padrão o Go vai reclamar do pacote pois possui oiutras dependencias que não estarão explicitamente aqui
	// Então ele pede um comentário de justificativa e tb um "apelido"
	_ "github.com/go-sql-driver/mysql"
)

func MySqlConnection() (*sql.DB, error) {

	urlConnection := config.StringConnection

	db, err := sql.Open("mysql", urlConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
