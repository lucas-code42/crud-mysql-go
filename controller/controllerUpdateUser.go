package controller

import (
	"crud-api-mysql/db"
	"crud-api-mysql/models"
	"database/sql"
	"fmt"
	"net/http"
)

func ControllerUpdateUser(r *http.Request) (models.User, error) {
	var user models.User

	body, _ := convertBodyToByte(r)
	convertBodyToStruct(body, &user)

	dataBase, err := db.MySqlConnection()
	if err != nil {
		fmt.Println("Erro de conezão com Banco", err)
		return models.User{}, err
	}
	defer dataBase.Close()

	ID, _ := getUrlParameters(r)
	alterUser(dataBase, user, ID)

	user.Id = uint(ID)

	return user, nil
}

// alterUser altera um usuário dentro do banco de dados
func alterUser(dataBase *sql.DB, user models.User, urlParameter int64) (int64, error) {
	var query string = "UPDATE crud_golang SET golang_name = ?, golang_email = ? WHERE golang_id = ?"

	stmt, err := dataBase.Prepare(query)
	if err != nil {
		fmt.Println("Erro no stateatment da query alterUser", err)
		return -1, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Email, urlParameter)
	if err != nil {
		fmt.Println("Erro na execução da query alterUser", err)
		return -1, err
	}

	row, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Erro ao pegar linhas afetadas", err)
		return -1, err
	}

	return row, nil
}
