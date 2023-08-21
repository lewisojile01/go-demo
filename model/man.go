package model

import "fmt"

type Man struct{}

type Book struct {
	ISDN    string
	Name    string
	Version string
}

func NewBook(name string) *Book {
	return &Book{
		Name: name,
	}
}
func (man *Man) Talk(word string) {
	startTalking(word)
}
func (man *Man) Walk(destination string) {
	fmt.Println("man is heading to ", destination)
}

func startTalking(word string) {
	fmt.Println("man said ", word)
}

func (book *Book) GetName() string {
	return book.Name
}
