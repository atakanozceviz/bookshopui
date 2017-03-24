package main

import (
	"github.com/atakanozceviz/bookshopui/controller"
)

func main() {
	port := "8080" //os.Getenv("PORT")
	controller.Start().Listen(":" + port)
}
