package main

import (
	"github.com/atakanozceviz/bookshopui/controller"
)

func main() {
	controller.Start().ListenLETSENCRYPT(":443")
}
