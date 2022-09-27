package controller

import (
	"crud-api-mysql/db"
	"crud-api-mysql/models"
	"database/sql"
	"fmt"
	"net/http"
)

// ControllerListUsers função principal do arquivo
func ControllerListUsers(r *http.Request) ([]models.User, error) {
	dataBase, err := db.MySqlConnection()
	if err != nil {
		fmt.Println("Erro de conexao com banco, ControllerListUsers", err)
		return []models.User{}, err
	}
	defer dataBase.Close()

	users, err := getUsersFromDb(dataBase)
	if err != nil {
		fmt.Println("Erro ao trazer o array de users", err)
		return []models.User{}, err
	}

	return users, nil
}

// getUsersFromDb faz a query para pegar todos usuários da tabela
func getUsersFromDb(dataBase *sql.DB) ([]models.User, error) {
	var users []models.User
	var user models.User

	var query string = "SELECT * FROM crud_golang"
	rows, err := dataBase.Query(query)
	if err != nil {
		fmt.Println("Erro na query do método getUsersFromDb", err)
		return []models.User{}, nil
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			fmt.Println("erroo aqui", err)
			return []models.User{}, nil
		}
		fmt.Println("*** -> ", user)
		users = append(users, user)
		fmt.Println(users)
	}

	return users, nil
}
