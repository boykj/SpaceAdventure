package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"
)

func main() {

	welcome()
	fmt.Print("Enter your name: ")
	var input string
	fmt.Scanln(&input)
	fmt.Println("Nice to meet you,", input)
	fmt.Println("My name is drangus, I'm an old friend of Kirk. Lets go on an adventure!")
	
	yesNo()

}

func getName() {
	for {
		fmt.Println("What's your name?")
		var input string
		fmt.Scanln(&input)
		fmt.Println(input)
	}
}

func welcome() {
	fmt.Println("Welcome to the Solar System!")
	fmt.Println("There are 9 planets to explore")
}

func yesNo() bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("[y/n]:")

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}
