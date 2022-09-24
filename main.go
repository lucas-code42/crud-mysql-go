package main

import (
	"crud-api-mysql/routers"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	textArt := figure.NewColorFigure("CRUD MySql GO", "", "cyan", true)
	textArt.Print()

	routers.StartRouters()
}

