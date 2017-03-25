package main

import (
	"os"

	"github.com/atakanozceviz/bookshopui/controller"
)

func main() {
	controller.Start().Listen(":" + os.Getenv("PORT"))
}
