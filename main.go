package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

var (
	words [16][2315]int
	radix [16][243]int
	table []byte
)

func lookup(guess, answer int) byte {
	return table[guess*2315+answer]
}

func solve(depth, start, end, cutoff int) (int, int) {
	if end-start == 1 {
		return 1, words[depth][start]
	}
	best := -1
	minimum := cutoff
	for i := start; i < end; i++ {
		maximum := 0
		for rad := 0; rad < 243; rad++ {
			radix[depth][rad] = 0
		}
		for j := start; j < end; j++ {
			radix[depth][lookup(words[depth][i], words[depth][j])] += 1
		}
		for rad := 1; rad < 243; rad++ {
			radix[depth][rad] += radix[depth][rad-1]
		}
		for j := start; j < end; j++ {
			score := lookup(words[depth][i], words[depth][j])
			radix[depth][score] -= 1
			words[depth+1][start+radix[depth][score]] = words[depth][j]
		}
		broken := false
		for rad := 242; rad > 0; rad-- {
			if start+radix[depth][rad-1] == start+radix[depth][rad] {
				continue
			}

			score, _ := solve(depth+1, start+radix[depth][rad-1], start+radix[depth][rad], minimum-1)
			if score > maximum {
				maximum = score
			}
			if maximum >= minimum {
				broken = true
				break
			}
		}
		if !broken {
			best = words[depth][i]
			minimum = maximum
		}
	}
	return 1 + minimum, best
}

func setup() {
	input, err := os.ReadFile("table.dat")
	if err != nil {
		panic(err)
	}
	table = make([]byte, 2315*2315)
	_, err = base64.StdEncoding.Decode(table, input)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 2315; i++ {
		words[0][i] = i
	}
}

func printWord(word int) {
	fmt.Println(word)
}

func readScore(score string) int {
	val, err := strconv.Atoi(score)
	if err != nil {
		panic(err)
	}
	return val
}

func makeMove(depth, move, score, total int) int {
	j := 0
	for i := 0; i < total; i++ {
		if int(lookup(move, words[depth][i])) == score {
			words[depth+1][j] = words[depth][i]
			j++
		}
	}
	return j
}

func main() {
	setup()
	depth := 0
	total := 2315
	move := 1636
	remainder := 5
	for i := 1; i < len(os.Args); i++ {
		score := readScore(os.Args[i])
		total = makeMove(depth, move, score, total)
		depth += 1
		remainder, move = solve(depth, 0, total, 15-depth)
	}
	fmt.Println(remainder)
	printWord(move)
}
