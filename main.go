package main

import (
	"bufio"
	"fmt"
	"os"
)

//AI core structure
type AI struct {
	name string
}

//Input functoin of AI.
//tentatively only for string value.
func (ai AI) Input(value string) {
	fmt.Println(value)
}

func (ai AI) greet() {
	fmt.Println("Hello, I am", ai.name, ".")
}

//NewAI creates a instance of AI structure
func NewAI(name string) *AI {
	ai := new(AI)
	ai.name = name
	return ai
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	//var ai *AI
	ai := NewAI("Yui")
	ai.greet()
	for stdin.Scan() {
		text := stdin.Text()
		fmt.Println(text)
		if text == "exit" {
			return
		}
	}
}
