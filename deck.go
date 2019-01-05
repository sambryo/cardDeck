package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

// function signature :-
// Go has ways of to do multiple return
// return type (deck, deck)
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// func WriteFile(filename string, data []byte, perm os.FileMode) error
// type converstion
// greeting := "Hi there!"
// fmt.Println([]byte(greeting))
// deck-> []string -> string-> []byte

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// 066 any one can read and write
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
