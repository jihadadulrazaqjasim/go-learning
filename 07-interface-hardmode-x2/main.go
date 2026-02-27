package main

import (
	"fmt"
	"os"
	"strings"
)

// Contracts
type Document interface {
	GetName() string
	GetSize() int
	Process() error
}

type Compressible interface {
	Compress() error
}

type Encryptable interface {
	Encrypt() error
}

type Logger interface {
	Log(message string)
}

// Implementations
type PDF struct {
	name string
	size int
}

func (p PDF) GetName() string {
	return p.name
}

func (p PDF) GetSize() int {
	return p.size
}

func (p PDF) Process() error {
	fmt.Printf("Processing PDF document (Size: %v KB) \n", p.size)
	return nil
}

func (p PDF) Encrypt() error {
	fmt.Println("Encrypting PDF document...")
	return nil
}

func (PDF) Compress() error {
	fmt.Println("Compressing PDF document...")
	return nil
}

type Word struct {
	name string
	size int
}

func (w Word) GetName() string {
	return w.name
}

func (w Word) Process() error {
	fmt.Printf("Processing Word document (Size: %v KB)\n", w.size)
	return nil
}

func (w Word) GetSize() int {
	return w.size
}

type Image struct {
	name       string
	size       int
	resolution string
}

func (img Image) GetName() string {
	return img.name
}

func (img Image) GetSize() int {
	return img.size
}

func (img Image) Process() error {
	fmt.Printf("Optimizing image (Resolution: %v, Size: %+v KB) \n", img.resolution, img.size)
	return nil
}

func (img Image) Compress() error {
	fmt.Println("Compressing image...")
	return nil
}

type ConsoleLogger struct{}

func (ConsoleLogger) Log(message string) {
	fmt.Println("[LOG] ", message)
}

type Processor struct {
	logger Logger
}

func (p Processor) Handle(doc Document) {

	p.logger.Log("Starting document: " + doc.GetName())

	//validation
	if doc.GetSize() < 0 {
		p.logger.Log("Error: document size cannot be negative")
		return
	}

	err := doc.Process()

	if err != nil {
		p.logger.Log("Error: " + err.Error())
		return
	}

	// //check type insertion for encrypt
	if enc, ok := doc.(Encryptable); ok {
		enc.Encrypt()
	} else {
		p.logger.Log("Encryption not supported")
	}

	// //check type insertion for compress
	if comp, ok := doc.(Compressible); ok {
		comp.Compress()
	} else {
		p.logger.Log("Compression not supported")
	}

	p.logger.Log("Finished document: " + doc.GetName())
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file name")
		return
	}

	fileName := os.Args[1]
	fileType := fileName[strings.LastIndex(fileName, "."):]

	consLog := ConsoleLogger{}
	process1 := Processor{logger: consLog}
	var doc Document
	switch fileType {
	case ".pdf":
		doc = PDF{name: fileName, size: 100}
	case ".docx", ".doc":
		doc = Word{name: fileName, size: 50}

	case ".png", ".jpeg", "jpg":
		doc = Image{name: fileName, size: 300, resolution: "1920x1080"}
	default:
		fmt.Println("Unsupported file type")
		return
	}
	process1.Handle(doc)

	fmt.Println()
	// Negative size
	wordNeg := Word{name: fileName, size: -1}
	process1.Handle(wordNeg)
}
