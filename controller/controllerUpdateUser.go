package controller

import (
	"crud-api-mysql/db"
	"crud-api-mysql/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func ControllerUpdateUser(r *http.Request) (string, error) {
	var user models.User
	
	body, _ := convertBodyToByte(r)

	dataBase, err := db.MySqlConnection()
	if err != nil {
		fmt.Println("Erro de conezão com Banco", err)
		return "", err
	}
	defer dataBase.Close()
	convertBodyToByte(r, )
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