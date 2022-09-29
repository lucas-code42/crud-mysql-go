package controller

import (
	"crud-api-mysql/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

// getUrlParameters pega o parametro da url e retorna um INT
func getUrlParameters(r *http.Request) (int64, error) {
	parameter := mux.Vars(r)
	ID, err := strconv.ParseInt(parameter["id"], 10, 64)
	if err != nil {
		fmt.Println("Erro ao converter parametro da url para INT", err)
		return -1, err
	}
	return ID, nil
}
