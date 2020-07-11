package main

import (

	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	var text string

	fmt.Println("Please introduce a string of characters")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	line := in.Text()

	text=line
	text = strings.ToLower(text)

	switch {
	case strings.HasPrefix(text,"i") && strings.Contains(text,"a") && strings.HasSuffix(text,"n"):
		fmt.Println("Found!")
	default:
		fmt.Println("Not Found!")
	}



}