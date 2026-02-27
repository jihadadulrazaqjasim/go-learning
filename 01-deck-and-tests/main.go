package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingHand := deal(cards, 10)

	hand.print()
	remainingHand.print()

	fmt.Println()

	cards.saveToFile("cards")

	loadedCards := newDeckFromFile("cards")
	loadedCards.shuffle()
	loadedCards.print()
}
