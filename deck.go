package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)
type deck []string

func (d deck) print() {
	for i, v := range d {
		fmt.Println(i, v)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(fileName string) error {
	deckBytes := []byte(d.toString())

	return os.WriteFile(fileName, deckBytes, 0644)
}

func (d deck) readFromFile(fileName string) deck {
	bytes, error := os.ReadFile(fileName)

	if error != nil {
		panic(error)
	}

	deckFromFile := strings.Split(string(bytes), ",")

	return deck(deckFromFile)
}

func newDeck() deck {
	newDeck := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			newDeck = append(newDeck, fmt.Sprintf("%s of %s", value, suit))
		}
	}

	return newDeck
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func shuffle(d deck) deck {
	rand.Seed(time.Now().UnixNano())
	
	for i, v := range d {
		temp := v
		swapIndex := rand.Intn(len(d) - 1)

		d[i] = d[swapIndex]
		d[swapIndex] = temp
	}

	return d
}
