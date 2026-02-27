package main

import "fmt"

type deck []string

func newDeck () deck {
	cards := deck{}
	cardSuits := []string {"Spades","Diamonds","Hearts"}
	cardValues := []string {"Ace","One","Two","Three","Four"}

	for _,suit := range cardSuits {
		for _,value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards

}

func (dInstance deck) print() {
	for i,card := range dInstance {
		fmt.Println(i,card)
	}
}