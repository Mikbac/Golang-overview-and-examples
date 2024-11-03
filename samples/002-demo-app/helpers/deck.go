package Helpers

type Deck []string

func NewDeck() Deck {
	cards := Deck{}

	typeA := []string{"AAA", "BBB", "CCC"}
	typeB := []string{"aaa", "bbb", "ccc"}

	for _, valueA := range typeA {
		for _, valueB := range typeB {
			cards = append(cards, valueA+" - "+valueB)
		}
	}

	return cards
}

func (d Deck) PrintDeck() {
	for i, e := range d {
		println(i, e)
	}
}
