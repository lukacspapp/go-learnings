package main

import "fmt"

func main() {

	// cards := newDeck()

	// cards.shuffle()

	// cards.print()

	num := []int{}

	for i := 1; i <= 10; i++ {
		num = append(num, i)
	}

	for n := range num {
		if n%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}

}
