package controller

import (
	applicationlog "crud-api-mysql/applicationLog"
	"crud-api-mysql/db"

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

	ContollerCreate(r)

	w.Write([]byte("CREATE"))
}
