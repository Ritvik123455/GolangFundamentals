package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"

	//cards := []string{newCard(), newCard()}

	cards2 := newDeck()
	// fmt.Println(cards2)
	// cards2.print()

	hand, remainingDeck := deal(cards2, 5)
	hand.print()
	fmt.Println("----")
	remainingDeck.print()
}
