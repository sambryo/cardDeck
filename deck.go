package main

import "fmt"

// Create a new type of 'deck which is a slice of strings

type deck []string

func newDeck() deck {

	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "eight", "nine", "ten",
		"Jack", "kings", "Queen"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// A recieve on a function
// Any variable of type "deck" now gets access to the "print" method
// why d -> convention one or two letter of the type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
