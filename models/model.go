package models

type User struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Bank struct {
	UserCredentials User    `json:"userCredentials"`
	Balance         float64 `json:"balance"`
	AccountId       string  `json:"accountId"`
}
