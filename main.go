package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)

	for stdin.Scan() {
		text := stdin.Text()
		fmt.Println(text)
		if text == "exit" {
			return
		}
	}
}
