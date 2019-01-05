package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard()}
	cards = append(cards, "six of spades")
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Aces"
}
