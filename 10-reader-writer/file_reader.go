package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(fileName string) {
	content, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	stat, errr := content.Stat()
	if errr != nil {
		fmt.Println("Error reading file stats:", errr)
		os.Exit(1)
	}

	fmt.Println("Name:", stat.Name(), " Size:", stat.Size(), " isDir:", stat.IsDir())

	if stat.Size() == 0 {
		fmt.Println("File is empty")
		return
	}

	fmt.Println("Content of file:")
	io.Copy(os.Stdout, content)
}
