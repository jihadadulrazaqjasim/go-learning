package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if(len(cards) != 16){
		t.Errorf("Expected deck length of 16, but got %v", len(cards))
	}

    expectedFirstCard := "Ace of Spades"
    if(cards[0] != expectedFirstCard) {
         t.Errorf("Expected first card to be %v but got %v",expectedFirstCard,cards[0])
    }

    expectedLastCard := "Four of Clubs"
    actualLastCard := cards[len(cards)-1]

    if(actualLastCard != expectedLastCard){
        t.Errorf("Expected last card to be %v but got %v",expectedLastCard,actualLastCard)
    }
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")

	 deck := newDeck()
	 deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	
	if(len(loadedDeck) != 16){
		t.Errorf("Expected length of loaded deck is 16 but got %v",len(loadedDeck))
	}
	os.Remove("_decktesting")
}