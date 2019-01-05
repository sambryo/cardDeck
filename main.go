package main

import "fmt"

func main() {
	// Slices in go
	cards := []string{newCard(), newCard()}
	cards = append(cards, "six of spades")

	// Iterate in to Cards
	for i, card := range cards {
		fmt.Println(i, card)
	}
	//fmt.Println(cards)
}

func newCard() string {
	return "Five of Aces"
}
