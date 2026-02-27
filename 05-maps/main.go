package main

import "fmt"

func main() {
	colors := make(map[string]string)
	colors["red"] = "#ff0000"
	colors["somecolor"] = "#ffbbbb"

	colors2 := map[string]string{"white": "#ffffff", "red": "#ff0000"}

	fmt.Println(colors, len(colors))
	fmt.Println(colors2, len(colors2))
}
