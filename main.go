package main

import (
	"fmt"
	"math/rand"
	"time"
)

var suites = [4]string{"D", "S", "C", "H"}
var ranks = [13]string{"A","2","3","4","5","6","7","8","9","10","J","Q","K"}

func createDesck() []string {
	deck := []string{}
	for _, s := range suites{
		for _, r := range ranks {
			deck = append(deck, s+r)
		}
	}
	return deck
}

func shuffleDeck(a []string) []string {
	x := time.Now().UnixNano()
	rand.Seed(x)
	fmt.Printf("x-seed:\n%+v\n",x)
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

func main() {
	deck := createDesck()
	fmt.Printf("Initial Deck: \n%+v\n", deck)
	deck = shuffleDeck(deck)
	fmt.Printf("Shuffled desck:\n%+v\n", deck)
	hands := deck[:5]
	fmt.Printf("Hands:\n%+v\n",hands)
}
