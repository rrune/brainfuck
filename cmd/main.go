package main

import (
	"log"
	"os"

	"github.com/rrune/brainfuck/interpreter"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Missing argument <file>")
	}
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("File not found")
	}
	file := string(f)

	it := interpreter.New()
	it.Run(file)
}
