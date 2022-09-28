package main

import (
	"os"
	"tetrimino"
)

func main() {
	if len(os.Args[1:]) != 1 {
		os.Exit(0)
	}
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		os.Exit(0)
	}
	tetrimino.PlayTetris(string(file))
}
