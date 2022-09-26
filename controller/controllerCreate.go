package controller

import (
	"crud-api-mysql/db"
	"crud-api-mysql/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ContollerCreate funcao princiapal do arquivo
func ContollerCreate(r *http.Request) (models.User, error) {
	var user models.User

	body, _ := convertBodyToByte(r)
	convertBodyToStruct(body, &user)

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

// convertBodyToByte converte o corpo de uma req em slice de bytes
func convertBodyToByte(r *http.Request) ([]byte, error) {
	newBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("convertBodyToByte ", err)
		return nil, err
	}
	return newBody, nil
}

// convertBodyToStruct recebe o corpo de uma requisicao em bytes e modela na referencia de um usuario
func convertBodyToStruct(body []byte, user *models.User) error {
	if err := json.Unmarshal(body, &user); err != nil {
		fmt.Println("convertBodyToStruct error ", err)
		return err
	}
	return nil
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
