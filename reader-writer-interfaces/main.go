package main

import (
	"fmt"
	// "io"
	// "net/http"
	// "os"
)

type logWriter struct{}

type shape interface {
	area() int
}

type square struct {
	sideLength int
}

type rectangular struct {
	length int
	width  int
}

func (r rectangular) area() int {
	return r.length * r.width
}

func (s square) area() int {
	return s.sideLength * s.sideLength
}

func main() {
	//http call to google.com
	// resp, err := http.Get("http://www.google.com")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }

	// lw := logWriter{}

	// io.Copy(lw, resp.Body)

	square1 := square{sideLength: 2}
	printArea(square1)

	rect1 := rectangular{length: 2, width: 3}
	printArea(rect1)

}

func printArea(s shape) {
	fmt.Printf("Type: %T, Size of the shape is: %d\n", s, s.area())
}

func (logWriter) Write(bs []byte) (n int, err error) {
	if len(bs) < 1 {
		return 1, fmt.Errorf("Size is empty baby")
	}
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))

	return len(bs), nil
}
