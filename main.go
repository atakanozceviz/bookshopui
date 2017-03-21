package main

import (
	"os"

	"github.com/atakanozceviz/bookshopui/controller"
)

func main() {
	os.Setenv("PORT", "3001")
	port := "3001"
	controller.Start().Listen(":" + port)
}
