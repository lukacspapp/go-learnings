package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello There"
}

func (spanishBot) getGreeting() string {
	return "Hola Senor"
}

// If there is any other type that has a getGreeting function that returns a string
// assossiated with it then that type is now a member ot bot interface

// As Spanish and English bots are the member of the bot interface they can call the printGreeting()

// ! Interfaces are contract to help us manage types
