package controller

import (
	"crud-api-mysql/db"
	"crud-api-mysql/models"
	"crud-api-mysql/util"
	"database/sql"
	"fmt"
	"net/http"
)

// ContollerCreate funcao princiapal do arquivo
func ContollerCreate(r *http.Request) (models.User, error) {
	var user models.User

	body, _ := util.ConvertBodyToByte(r)
	util.ConvertBodyToStruct(body, &user)

	dataBase, err := db.MySqlConnection()
	if err != nil {
		return models.User{}, err
	}
	defer dataBase.Close()

	err = executeQuery(dataBase, user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil

}

// executeQuery prepara e executa uma query
func executeQuery(dataBase *sql.DB, user models.User) error {
	stmt, err := dataBase.Prepare("INSERT INTO crud_golang(golang_name, golang_email) values(?, ?)")
	if err != nil {
		fmt.Println("executeQuery error ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email)
	if err != nil {
		fmt.Println("executeQuery error", err)
		return err
	}
	return nil

}
