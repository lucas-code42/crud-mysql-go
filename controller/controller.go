package controller

import (
	applicationlog "crud-api-mysql/applicationLog"
	"crud-api-mysql/db"
	"encoding/json"

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
		w.Write([]byte("Fail to create new user"))
		return
	}

	msg := map[string]string{"message": "user created"}
	returnStmt, err := json.Marshal(msg)
	if err != nil {
		w.Write([]byte("Fail to create new user"))
		return
	}

	w.Write(returnStmt)
}
