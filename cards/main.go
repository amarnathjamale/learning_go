package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.saveToFile("all_cards")
	hand, remainingCards := deal(cards, 3)
	hand.print()
	remainingCards.print()
	fmt.Println(hand.toString())
	readFromFile("all_cards").print()
	cards.print()
}
