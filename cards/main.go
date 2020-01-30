package main

func main() {

	cards := newDeck()
	cards.saveToFile("my_file")

	retrieved_cards := savedDeck("my_file")

	retrieved_cards.print()
}
