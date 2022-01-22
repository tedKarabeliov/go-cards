package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	
	d := newDeck()

	d.print()

	d = shuffle(d)

	fmt.Println("-----------------------------------------")

	d.print()
}