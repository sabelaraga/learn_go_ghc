// TODO: make it a function of the type.

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
	m := Message{text:"Welcome to GHC!"}
	m.print()
}
