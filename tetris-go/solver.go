package tetrimino

/*
Removes the tetrimino and puts back dots.
*/

func RemoveBlock(tetris tetr, block []rune, x, y int) {
	resetX := x
	for _, c := range block {
		if c == '\n' {
			x = resetX
			y++
		} else {
			if c != '.' {
				tetris.dotmap[y][x] = '.'
			}
			x++
		}
	}
}

/*
Inserts the tetrimino to given spot
*/

func InsertBlock(tetris tetr, block []rune, x, y int) {
	resetX := x
	for _, c := range block {
		if c == '\n' {
			x = resetX
			y++
		} else {
			if c != '.' {
				tetris.dotmap[y][x] = c
			}
			x++
		}
	}
}

/*
Checks if the tetrimino fits in.
If there's already placed something
it returns false.
*/

func FindBlock(tetris tetr, block []rune, x, y int) bool {
	resetX := x
	for _, c := range block {
		if c == '\n' {
			x = resetX
			y++
		} else {
			if x >= tetris.size || y >= tetris.size {
				return false
			}
			if c != '.' && tetris.dotmap[y][x] != '.' {
				return false
			}
			x++
		}
	}
	return true
}

/*
Tries to find a spot for the first tetrimino,
if successful it inserts the tetrimino and
tries to place the next tetrimino.
If unsuccessful it can't place the next tetrimino
it removes the current one.
*/

func PlaceIt(tetris tetr, blocks [][]rune, index int) bool {
	for y := 0; y < tetris.size; y++ {
		for x := 0; x < tetris.size; x++ {
			if FindBlock(tetris, tetris.shapes[index], x, y) == true {
				InsertBlock(tetris, blocks[index], x, y)
				if index == tetris.bc-1 || PlaceIt(tetris, blocks, index+1) == true {
					return true
				} else {
					RemoveBlock(tetris, blocks[index], x, y)
				}
			}
		}
	}
	return false
}

/*
Creates a map that contains just dots.
Tries to place tetriminos to the map,
if unsuccessful creates a larger map and
tries again. If successful it prints
the result.
*/

func Solver(tetris tetr, size int) {
	tetris.dotmap = CreateMap(tetris, size)
	tetris.size = size
	if PlaceIt(tetris, tetris.shapes, 0) == false {
		Solver(tetris, tetris.size+1)
	} else {
		PrintMap(tetris)
	}
}
