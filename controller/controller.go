package controller

import (
	applicationlog "crud-api-mysql/applicationLog"
	"crud-api-mysql/db"

	"net/http"
)

func HelloWeb(w http.ResponseWriter, r *http.Request) {
	applicationlog.LogHelloWeb()

	dataBase, err := db.ConnectionMySql()
	if err != nil {
		return
	}
	defer dataBase.Close()
	w.Write([]byte("Hello web!"))
}

func Create(w http.ResponseWriter, r *http.Request) {
	applicationlog.LogCreate()

	ContollerCreate(r)

	w.Write([]byte("CREATE"))
}
