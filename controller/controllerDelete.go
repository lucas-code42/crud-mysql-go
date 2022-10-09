package controller

import (
	"crud-api-mysql/db"
	"crud-api-mysql/util"
	"database/sql"
	"fmt"
	"net/http"
)

// ControllerDelete função principal do arquivo
func ControllerDelete(r *http.Request) (int64, error) {
	dataBase, err := db.MySqlConnection()
	if err != nil {
		fmt.Println("Erro de conexão com Banco de Dados", err)
		return -1, err
	}
	defer dataBase.Close()

	ID, _ := util.GetUrlParameters(r)

	deletedId, err := deleteOnDataBase(dataBase, ID)
	if err != nil {
		return -1, err
	}

	return deletedId, nil
}

// deleteOnDataBase cria e executa uma query para deletar dados do banco
func deleteOnDataBase(dataBase *sql.DB, id int64) (int64, error) {
	var query string = "DELETE FROM crud_golang WHERE(golang_id = ?)"
	stmt, err := dataBase.Prepare(query)
	if err != nil {
		fmt.Println("Erro ao preparar steatment", err)
		return -1, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		fmt.Println("Erro ao executar query de delete", err)
		return -1, err
	}

	return result.LastInsertId()
}
