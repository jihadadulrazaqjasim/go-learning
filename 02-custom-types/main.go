package main

import (
	"fmt"
	"strings"
)

type book string

func (b book) printTitle() {
	fmt.Println(b)
}

func (b book) countLetters() int {
	return len(b)
}

func (b book) titleToUpperCase() string {
	return strings.ToUpper(string(b))
}

type color string

func (c color) describe(description string) string {
	return string(c) + " " + description
}

func main() {
	b := book("First things First")
	b.printTitle()
	fmt.Println("Letters:", b.countLetters())
	fmt.Println(b.titleToUpperCase())

	c := color("Red")
	fmt.Println(c.describe("is an awesome color"))

	deck := newDeck()
	deck.print()
}
