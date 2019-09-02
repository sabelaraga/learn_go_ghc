// TODO: add a function that creates a Message.

package main

import (
	"fmt"
)

type Message struct{
	text string
}

func (m Message) print() {
	fmt.Println(m.text)
}


func main() {
	m := NewMessage("Welcome to GHC!")
	m.print()
}