// TODO: add a function that prints the message.

package main

import (
	"fmt"
)

type Message struct{
	text string
}

func print (m Message) {
	fmt.Println(m.text)
}

func main() {
	m := Message{text:"Welcome to GHC!"}
	print(m)
}
