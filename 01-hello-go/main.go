package main

// import "fmt"

func main(){
// 	cards := newDeck()
//     hand,remainingHand:= deal(cards, 10)

// 	hand.print()
// 	remainingHand.print()

// 	 c := color("Red")
 
//    fmt.Println(c.describe("is an awesome color"))

// cards := newDeck
// cards.saveToFile("cards")

// bytes := []byte (cards.toString())
// fmt.Println(bytes)

// cards := newDeckFromFile("myyy")
// cards.print()

cards := newDeck()
cards.shuffle()
cards.print()

}

type color string
 
func (c color) describe(description string) (string) {
   return string(c) + " " + description
}