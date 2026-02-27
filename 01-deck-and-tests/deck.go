package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	suits := []string{"Spades","Diamonds","Hearts","Clubs"}
	values := []string{"Ace","Two","Three","Four"}

	for _,suit := range suits {
		for _,value := range values{
		cards = append(cards, value + " of "+ suit)
		}
	}
	
	return cards

}

func (d deck) print (){
	for i,card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string{

	return 	strings.Join([]string(d),",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename,[]byte(d.toString()),0666)
}

func newDeckFromFile(filename string) deck {

byteSlices,err := os.ReadFile(filename)

if(err != nil){
	fmt.Println("Error while reading file: ", err)
	os.Exit(1)
}

return deck(strings.Split(string(byteSlices),","))
}

func (d deck) shuffle (){
	for i := range d {
		newIndex := rand.IntN(len(d)-1)
		d[i],d[newIndex] = d[newIndex], d[i]
	}

}