package controller

import (
	"net/http"
)

func ContollerCreate(r *http.Request) ([]byte, error) {
	return []byte("teste"), nil

}
