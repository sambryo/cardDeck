package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// option 1 - log the error & return a call to newDeck.
		// option 2 - log the error & entirely quit the program.
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

// shuffle a deck of card
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // to generate a different seed fromr random generator.
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
