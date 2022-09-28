package tetrimino

func Match(s string, toFind string) bool {
	al, bl := len(s), len(toFind)
	a := []rune(s)
	b := []rune(toFind)
	if al == 0 || bl == 0 {
		return false
	}
	for i := 0; i < al; i++ {
		if a[i] == b[0] {
			temp := a[i : i+bl]
			if string(temp) == toFind {
				return true
			}
		}
	}
	return false
}

func FindShapes(block string) string {
	if Match(block, "##..##") {
		return "##\n##"
	}
	if Match(block, ".##.##") {
		return ".##\n##."
	}
	if Match(block, "#...##...#") {
		return ("#.\n##\n.#")
	}
	if Match(block, "###..#") {
		return ("###\n.#.")
	}
	if Match(block, ".#..##...#") {
		return (".#\n##\n.#")
	}
	if Match(block, ".#..###") {
		return (".#.\n###")
	}
	if Match(block, "#...##..#") {
		return ("#.\n##\n#.")
	}
	if Match(block, "##...##") {
		return "##.\n.##"
	}
	if Match(block, ".#..##..#") {
		return (".#\n##\n#.")
	}
	if Match(block, "####") {
		return "####"
	}
	if Match(block, "#...#...#...#") {
		return ("#\n#\n#\n#\n")
	}
	if Match(block, "#...#..##") {
		return (".#\n.#\n##")
	}
	if Match(block, "#...###") {
		return ("#..\n###")
	}
	if Match(block, "##..#...#") {
		return ("##\n#.\n#.")
	}
	if Match(block, "###...#") {
		return ("###\n..#")
	}
	if Match(block, "#...#...##") {
		return ("#.\n#.\n##")
	}
	if Match(block, "###.#") {
		return ("###\n#..")
	}
	if Match(block, "##...#...#") {
		return ("##\n.#\n.#")
	}
	if Match(block, "..#.###") {
		return ("..#\n###")
	}
	return ""
}

/*
Makes each block into smaller string
and gives it a correct letter
*/

func Identify(tetris tetr) [][]rune {
	tetris.char = 'A'
	var shapes []string
	for _, block := range tetris.block {
		shapes = append(shapes, FindShapes(block))
	}
	for _, shape := range shapes {
		var arr []rune
		for _, c := range shape {
			if c == '#' {
				c = rune(tetris.char)
				arr = append(arr, c)
			} else {
				arr = append(arr, c)
			}
		}
		tetris.shapes = append(tetris.shapes, arr)
		tetris.char++
	}
	return tetris.shapes
}
