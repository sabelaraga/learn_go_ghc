// TODO: add a text field to the Message, initialize it, and print it.

package main

import (
	"fmt"
)

type Message struct{
}

func main() {
	m := Message{}
	fmt.Println(m.text)
}
