// package main

// // import "fmt"

// func main() {
//   // theNumber := getNumber()
//   // var anotherNumber int  = 123;

//   // fmt.Println(theNumber)
//   // fmt.Println(anotherNumber)

//   // cards := [] string {
//   //   "Frist Card", newCard(),
//   // }

//   cards := deck {
//     "Frist Card", newCard(),
//   }

// //   for index,card := range cards {
// //     fmt.Println(index,card)
// // }

// cards.print();

//   // updatedCards := append(cards, "ahahah")

//   // for i,card := range updatedCards {
//   //   fmt.Println(i, card);
//   // }

// }

// // func getNumber () int {
// // 	return 1;
// // }

// func newCard() string {
//   return "New Card";
// }

package main

import ("fmt"; "strings")

type book string

func (b book) printTitle(){
  fmt.Println(b)
}

func (b book) countLetters() int {
   return len(b)
}

func (b book) titleToUpperCase() string{
  return strings.ToUpper(string(b));
}

func main(){
  // var b book = "First things First"

    deck := newDeck()

    deck.print();


  // b.printTitle();
  // fmt.Println("Letters: ", b.countLetters())
  // fmt.Println(b.titleToUpperCase())

}