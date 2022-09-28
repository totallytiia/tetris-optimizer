package tetrimino

import (
	"fmt"
	"os"
)

func ErrExit() {
	fmt.Print("ERROR")
	os.Exit(0)
}

/*
Counts how many tetriminos is
in the text file.
*/

func BlockCounter(file string) int {
	count := 0
	for i := 0; i < len(file); {
		if file[i] == '\n' && i%5 != (4+count)%5 && i != len(file)-1 {
			ErrExit()
		}
		i++
		if i == 20+21*count || i == len(file)-2 {
			count++
		}
	}
	if 21*count-2 != len(file) {
		ErrExit()
	}
	return count
}

/*
Checks that each # is connected together
*/

func CheckConnect(file string, i int) int {
	check := 0
	if i+1 < len(file) && file[i+1] == '#' {
		check++
	}
	if i >= 1 && file[i-1] == '#' {
		check++
	}
	if i >= 4 && file[i-4] == '#' {
		check++
	}
	if i+4 < len(file) && file[i+4] == '#' {
		check++
	}
	return (check)
}

/*
Checks that the blocks have only . and # and
that there's correct amount of them.
Each block needs to have 12 dots and
4 hashes. Function also checks that
the shapes are valid.
*/

func Validate(tetris tetr) {
	dot, hash, connect := 0, 0, 0
	for _, s := range tetris.block {
		connect = 0
		for i, c := range s {
			if c == '.' {
				dot++
			}
			if c == '#' {
				connect += CheckConnect(s, i)
				hash++
			}
		}
		if connect == 6 || connect == 8 {
			continue
		} else {
			ErrExit()
		}
	}
	if dot != 12*tetris.bc && hash != 4*tetris.bc {
		ErrExit()
	}
}
