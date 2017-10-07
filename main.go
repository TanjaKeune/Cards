package main

func main() {
	//	cards := newDeck()
	//	hand, remainingDeck := deal(cards, 5)

	//	hand.print()
	//	fmt.Println("______________________")
	//	remainingDeck.print()
	//	cards.print()

	cards := newDeck()
	cards.saveToFile("myCards")

	cards.shuffle()
	cards.print()
}
