package controller

import "go-demo/interfaces"

func Talk(human interfaces.IPerson, word string) {
	human.Talk(word)
}
func Walk(human interfaces.IPerson, destination string) {
	human.Walk(destination)
}
