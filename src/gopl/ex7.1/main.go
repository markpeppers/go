package main

import (
	"bufio"
	"fmt"
)

// WordCounter counts words
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	place := 0
	count := 0
	for place < len(p) {
		advance, _, _ := bufio.ScanWords(p[place:], true)
		count++
		place += advance
	}
	*c = WordCounter(count)
	return count, nil
}

// LineCounter counts lines
type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	place := 0
	count := 0
	for place < len(p) {
		advance, _, _ := bufio.ScanLines(p[place:], true)
		count++
		place += advance
	}
	*c = LineCounter(count)
	return count, nil
}

func main() {
	var c WordCounter
	c.Write([]byte("what up world\nHere we go"))
	fmt.Println(c)

	c = 0
	foo := "Thief"
	fmt.Fprintf(&c, "Hail to the %s\n", foo)
	fmt.Println(c)

	var d LineCounter
	fmt.Fprintf(&d, "Fuck you\nI know who %s is\nNow beat it", foo)
	fmt.Println(d)
}
