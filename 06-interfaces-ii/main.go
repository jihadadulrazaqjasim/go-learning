package main

import (
	"fmt"
)

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (ConsoleLogger) Log(message string) {
	fmt.Println("ConsoleLogger: ", message)
}

type FileLogger struct{}

func (FileLogger) Log(message string) {
	fmt.Println("FileLogger: ", message)
}

type Processor struct {
	logger Logger
}

type Document interface {
	GetName() string
	Process() error
}

type PDF struct {
	name string
}

func (p PDF) GetName() string {
	return p.name
}

func (p PDF) Process() error {
	fmt.Println("Processing PDF document..")
	return nil
}

func (p Processor) Handle(doc Document) {
	p.logger.Log("Starting the document processing: " + doc.GetName())
	err := doc.Process()

	if err != nil {
		p.logger.Log("Error: " + err.Error())
		return
	}

	p.logger.Log("End of document processing: " + doc.GetName())
}

func main() {
	conLog := ConsoleLogger{}
	processor1 := Processor{logger: conLog}
	pdf := PDF{name: "myfile.pdf"}
	processor1.Handle(pdf)

	processor1.logger.Log("\n\n")

	fileLogger := FileLogger{}
	processor2 := Processor{logger: fileLogger}
	processor2.Handle(pdf)
}
