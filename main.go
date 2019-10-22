package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"
)

func main() {

	//Display the welcome message
	welcome()
	
	//Early dictionary creation for hard coding the planetary values
	//Creates a map of strings
	p := make(map[string][]string)
	p["Mercury"] = []string{"Name: ", "Mercury", " || Description: ", "A very hot planet, closest to the sun"}
	p["Venus"] = []string{"Venus", "It's very cloudy here!"}
	p["Earth"] = []string{"Earth", "There is something very familiar about this planet"}
	p["Mars"] = []string{"Mars", "Known as the red planet."}
	p["Jupiter"] = []string{"Jupiter", "A gas giant, with a noticeable red spot."}
	p["Saturn"] = []string{"Saturn", "This planet has beautiful rings around it."}
	p["Uranus"] = []string{"Uranus", "Strangely, this planet rotates around on its side."}
	p["Neptune"] = []string{"Neptune", "A very cold planet, furthest from the sun."}
	p["Pluto"] = []string{"Pluto", "I don't care what they say - it's a planet."}
	
	//fmt.Println(p["Mc"])
	//fmt.Println(p["Vn"])
	
	for k, v := range p {
		fmt.Println("Planet: ", k, "Description: ", v)
	}
	
	//Prompt user to enter their name and echo
	
	fmt.Println("What's your name?")
	var input string
	fmt.Scanln(&input)
	fmt.Println("Nice to meet you,", input)
	fmt.Println("My name is drangus, I'm an old friend of Kirk. Lets go on an adventure!")
	
	//Prompt for user to input yes or no for space planet selection
	
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

func pSelect(map) {
		
		fmt.Println("Please select a planet from our solar system")
		fmt.Print("Planet name: ")
		var pInput string
		fmt.Scanln(&pInput)

		for k, v := range map {
			if pInput == k {
			fmt.Println(k, v)
			
			}
		}
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
		
			//If yes, display prompt to user to enter a planet name and print corresponding sentences
			
			fmt.Println("Okay, select a planet from our solar system")
			fmt.Print("Planet name: ")
			var input string
			fmt.Scanln(&input)
			pSelect()
			return true
			
		} else if response == "n" || response == "no" {
			
			//If no, return false and randomly choose a planet
			
			fmt.Println("Fair enough, I'll choose for you")
			return false
		} else {
			fmt.Println("Please enter a valid answer Y/N, or Yes / No")
		}
	}

}



