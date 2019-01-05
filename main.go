package main

func main() {
	// Slices in go
	cards := newDeck()
	//cards.print()
	//fmt.Println(cards)

	hand, remaingDeck := deal(cards, 5)
	hand.print()
	print("************")
	remaingDeck.print()
}
