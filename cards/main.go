package main

func main() {
	cards := deck{newCard(), "Ace of Diamonds"} // cards is of type deck
	cards = append(cards, "Six of Spades")

	cards.print() // bc cards is of type deck it can call print method
}

func newCard() string {
	return "Five of Diamonds"
}
