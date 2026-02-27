package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]
	content, err := os.Open(fileName)
	stat, errr := content.Stat()
	fmt.Println("Name: ", stat.Name(), " Size: ", stat.Size(), " isDir: ", stat.IsDir())

	fmt.Println("Content of file: ")
	if err != nil || errr != nil || stat.Size() == 0 {
		fmt.Println("Error opening file: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, content)

}
