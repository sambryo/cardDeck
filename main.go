package main

func main() {
	// Slices in go
	cards := deck{newCard(), newCard()}
	cards = append(cards, "six of spades")

	cards.print()
	//fmt.Println(cards)
}

func newCard() string {
	return "Five of Aces"
}
