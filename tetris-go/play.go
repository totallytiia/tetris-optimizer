package tetrimino

import (
	"strings"
)

type tetr struct {
	shapes [][]rune
	dotmap [][]rune
	block  []string
	bc     int
	size   int
	char   byte
}

/*
Splits the string into correct size (len == 16),
removes all newlines, and each string to tetris.block
*/

func PlayTetris(file string) {
	var tetris tetr
	tetris.bc = BlockCounter(file)
	temp := strings.ReplaceAll(file, "\n", "")
	for i := 0; i < len(temp); i += 16 {
		end := i + 16
		tetris.block = append(tetris.block, temp[i:end])
	}
	Validate(tetris)
	tetris.shapes = Identify(tetris)
	Solver(tetris, MapSize(tetris.bc*4))
}
