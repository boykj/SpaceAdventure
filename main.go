package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	welcome()
	getName()

}

func getName() {
	for {
		reader := bufio.NewReader(os.Stdin)
		var name string
		fmt.Println("What is your name?")
		name, _ := reader.ReadString('\n')
		name = strings.Replace(name, "\r\n", "", -1)
		fmt.Println("Your name is ", name)
	}
}

func welcome() {
	fmt.Println("Welcome to the Solar System!")
	fmt.Println("There are 9 planets to explore")
}
