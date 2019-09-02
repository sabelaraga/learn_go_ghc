// TODO: add a text field to the Message, initialize it, and print it.

package main

import (
	"fmt"
)

type Message struct{
	text string
}

func main() {
	m := Message{text:"Welcome to GHC!"}
	fmt.Println(m.text)
}
