package main

import (
	"fmt"
	"go-demo/controller"
	"go-demo/model"
)

func main() {
	man, woman := model.Man{}, model.Woman{}
	controller.Talk(&man, "coding is interesting")
	controller.Talk(&woman, "coding is interesting")
	controller.Walk(&man, "Rio")
	controller.Walk(&woman, "Trinidad")

	book := model.NewBook("Golang Complete Guide")

	fmt.Printf("%s\n",book.GetName())
}
