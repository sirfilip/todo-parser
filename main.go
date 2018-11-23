package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("TODO.txt")
	if err != nil {
		log.Fatal(err)
	}
	parser := NewParser(file)
	todos, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Got list of todos: %#v", todos)
}
