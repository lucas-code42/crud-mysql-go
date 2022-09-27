package controller

import (
	applicationlog "crud-api-mysql/applicationLog"
	"crud-api-mysql/db"
	"encoding/json"
	"fmt"

	"net/http"
)

func HelloWeb(w http.ResponseWriter, r *http.Request) {
	applicationlog.LogHelloWeb()

	dataBase, err := db.MySqlConnection()
	if err != nil {
		return
	}
	defer dataBase.Close()
	w.Write([]byte("Hello web!"))
}

// Create cria um usuário no banco
func Create(w http.ResponseWriter, r *http.Request) {
	applicationlog.LogCreate()

	_, err := ContollerCreate(r)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Fail to create new user"))
		return
	}

	msg := map[string]string{"message": "user created"}
	returnStmt, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Fail to return msg user"))
		return
	}

	w.Write(returnStmt)
}

// Delete deleta um usuário a partir de um ID
func Delete(w http.ResponseWriter, r *http.Request) {
	applicationlog.LogDelete()
	_, err := ControllerDelete(r)
	if err != nil {
		w.Write([]byte("Fail to create new user"))
		return
	}

	msg := map[string]string{"message": "user deleted"}
	returnStmt, err := json.Marshal(msg)
	if err != nil {
		w.Write([]byte("Fail to delete user"))
		return
	}

	w.Write(returnStmt)
}

// ListAllUsers lista todos os usuarios da tabela do banco de dados
func ListAllUsers(w http.ResponseWriter, r *http.Request) {
	applicationlog.LogListAllUsers()
	users, err := ControllerListUsers(r)
	if err != nil {
		w.Write([]byte("Fail to list users"))
		return
	}

	returnStmt, err := json.Marshal(users)
	if err != nil {
		w.Write([]byte("Fail to delete user"))
		return
	}

	w.Write(returnStmt)
}
