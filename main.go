package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"
	//"math/rand"
)

func main() {

	//Display the welcome message and get user name
	welcome()
	getName()
	
	//Early dictionary creation for hard coding the planetary values
	//Creates a map of strings
	p := make(map[string][]string)
	p["Mercury"] = []string{"Description:", "A very hot planet, closest to the sun"}
	p["Venus"] = []string{"Description:", "It's very cloudy here!"}
	p["Earth"] = []string{"Description:", "There is something very familiar about this planet"}
	p["Mars"] = []string{"Description:", "Known as the red planet."}
	p["Jupiter"] = []string{"Description:", "A gas giant, with a noticeable red spot."}
	p["Saturn"] = []string{"Description:", "This planet has beautiful rings around it."}
	p["Uranus"] = []string{"Description:", "Strangely, this planet rotates around on its side."}
	p["Neptune"] = []string{"Description:", "A very cold planet, furthest from the sun."}
	p["Pluto"] = []string{"Description:", "I don't care what they say - it's a planet."}
	
	//Displays the contents of planet map to the console
	fmt.Println("Have a look at the planets within our solar system:")
	fmt.Println()
	for k, v := range p {
		fmt.Println("Planet: ", k, "||", v)
	}
	
	//Prompt for user to input yes or no for space planet selection
	
	//yesNo()
	
	//If yes, prompt user to enter planet name with proper puncuation
	//Will print travelling to, and the key (planet name) then value (description)
	if yesNo() == true {
		fmt.Println("Please select a planet from our solar system (Use capitals || Case sensitive in this build)")
		fmt.Print("Planet name: ")
		var pInput string
		fmt.Scanln(&pInput)
		
		for k, v := range p {
			if k == pInput {
			fmt.Print("Travelling to -> ")
			fmt.Println(k, v)
			
			}
		}
	//If no, randomly select a planet from the map 
	} else {
		for k := range p {
			fmt.Println(k)
		}
		
	}
}

//Prompts user to input their name and echo accordinly - single strings at the moment
func getName() {

	fmt.Println("What's your name?")
	var input string
	fmt.Scanln(&input)
	fmt.Println("Nice to meet you,", input)
	fmt.Println("My name is drangus, I'm an old friend of Kirk. Lets go on an adventure!")
	fmt.Println()
	
}

func welcome() {
	fmt.Println("Welcome to the Solar System!")
	fmt.Println("There are 9 planets to explore")
}

func yesNo() bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("If you want to choose a planet, say yes. Otherwise, say no")
		fmt.Printf("[y/n]:")

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
		
			//If yes, display prompt to user to enter a planet name and print corresponding sentences
			
			//fmt.Println("Okay, select a planet from our solar system")
			//fmt.Print("Planet name: ")
			//var input string
			//fmt.Scanln(&input)
			//pSelect()
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

//func randKey(m map[int]int) int { 		-> Potential header format for a map argument function?
//func randKey(m map[string]string) string {
    //for k := range m {
        //return k
    //}
    //return k



