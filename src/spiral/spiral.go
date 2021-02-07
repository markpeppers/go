package main

import "fmt"

func main() {
	s := spiral(4)
	for line := range s {
		fmt.Println(s[line])
	}
}

// s[y][x]
func spiral(n int) [][]int {
	s := make([][]int, n)
	for i := range s {
		s[i] = make([]int, n)
	}
	x, y := 0, 0
	dir := 0 // 0 right 1 down 2 left 3 up
	xmax, ymax := n-1, n-1
	xmin, ymin := 0, 1

	for i := 0; i < n*n; i++ {
		s[y][x] = i + 1
		switch dir {
		case 0:
			if x == xmax {
				y++
				dir = 1
				xmax--
			} else {
				x++
			}
		case 1:
			if y == ymax {
				x--
				dir = 2
				ymax--
			} else {
				y++
			}
		case 2:
			if x == xmin {
				y--
				dir = 3
				xmin++
			} else {
				x--
			}
		case 3:
			if y == ymin {
				x++
				dir = 0
				ymin++
			} else {
				y--
			}
		}
	}
	return s
}
