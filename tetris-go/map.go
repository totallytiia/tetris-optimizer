package tetrimino

import (
	"fmt"
)

func MapSize(nb int) int {
	var size int
	for i := 2; i*i < nb; i++ {
		size = i
	}
	return size + 1
}

func GetRows(size int) []rune {
	var getRow []rune
	for i := 0; i < size; i++ {
		getRow = append(getRow, '.')
	}
	return getRow
}

func CreateMap(tetris tetr, size int) [][]rune {
	var dotmap [][]rune
	tetris.size = size
	for i := 0; i < size; i++ {
		dotmap = append(dotmap, GetRows(size))
	}
	return dotmap
}

func PrintMap(tetris tetr) {
	for _, line := range tetris.dotmap {
		fmt.Println(string(line))
	}
}
