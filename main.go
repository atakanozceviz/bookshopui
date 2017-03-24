package main

import (
	"os"

	"github.com/atakanozceviz/bookshopui/controller"
)

func main() {
	port := os.Getenv("PORT")
	controller.Start().Listen(":" + port)
}
