package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)
	hand2, remainingDeck1 := deal(remainingDeck, 5) //passed as arguments
	hand.print()                                    //used as receivers
	hand2.print()
	remainingDeck1.print()

	str := toString(hand)
	str1 := hand2.toString1()
	fmt.Println(str1)
	fmt.Println(str)

	cards.saveToFile("my_cards")
	newCards := newDeckfromFile("my_cards")
	newCards.print()

	newCards.shuffle()
	newCards.print()
}
