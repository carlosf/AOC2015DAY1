package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func p(str string) {
	fmt.Println(str)
}
func main() {

	p("Starting Day 1 2015 Advent of Code")

	f, err := os.Open("input.txt")
	check(err)

	defer f.Close()

	reader := bufio.NewReader(f)
	floor := 0
	iter := 0
	pos := 0
	for {
		char, err := reader.ReadByte()
		if err != nil {
			break
		}
		iter++
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
			if floor == -1 && pos == 0 {
				pos = iter
				fmt.Printf("basement at %d\n", pos)
			}
		}
		//fmt.Printf("total floor %d", floor)

		//fmt.Printf("%c", char)
	}
	fmt.Printf("total floor %d\n", floor)
}
