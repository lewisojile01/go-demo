package model

import "fmt"

type Woman struct{}

func (woman *Woman) Talk(word string) {
	fmt.Println("woman said ", word)
}
func (woman *Woman) Walk(destination string) {
	fmt.Println("woman is heading to ", destination)
}
