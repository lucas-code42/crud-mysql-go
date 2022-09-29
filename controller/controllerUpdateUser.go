package controller

import (
	"crud-api-mysql/db"
	"crud-api-mysql/models"
	"fmt"
	"net/http"
)

func ControllerUpdateUser(r *http.Request) (string, error) {
	var user models.User

	body, _ := convertBodyToByte(r)
	convertBodyToStruct(body, &user)

	dataBase, err := db.MySqlConnection()
	if err != nil {
		fmt.Println("Erro de conezão com Banco", err)
		return "", err
	}
	defer dataBase.Close()

	return "", nil
}
