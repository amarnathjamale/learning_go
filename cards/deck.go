package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardsSuits := []string{"Spades", "Club", "Hearts", "Diamonds"}
	cardsValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardsSuits {
		for _, value := range cardsValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0755)
}

func readFromFile(filename string) deck {
	_, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("File exists")
	}
	bs, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(bs), ",")
}

func (d deck) shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
