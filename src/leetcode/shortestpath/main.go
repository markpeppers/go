package main

import (
	"fmt"
)

type point struct {
	i int
	j int
}

type node struct {
	i int
	j int
	N int
}

func (n node) toPoint() point {
	return point{n.i, n.j}
}

func (n node) children(v *visited) []node {
	ch := make([]node, 0)
	for i := n.i - 1; i <= n.i+1; i++ {
		if i < 0 || i >= v.size {
			continue
		}
		for j := n.j - 1; j <= n.j+1; j++ {
			if j < 0 || j >= v.size {
				continue
			}
			if i == n.i && j == n.j {
				continue
			}
			newChild := node{i, j, n.N + 1}
			if newAdded := v.addNew(newChild.toPoint()); newAdded {
				ch = append(ch, newChild)
			}
		}
	}
	// fmt.Println("node:", n, " -- returning children: ", ch)
	return ch
}

type visited struct {
	size int
	list map[point]int
}

func newVisitedList(s int) visited {
	l := make(map[point]int)
	return visited{size: s, list: l}
}

// returns true if node was added (it's new to the list)
func (v *visited) addNew(p point) bool {
	if v.list[p] > 0 {
		return false
	}
	// Have a new member
	v.list[p] = 1
	return true
}

type queue struct {
	q []node
}

func newQueue() queue {
	newQ := make([]node, 0)
	return queue{q: newQ}
}

func (q *queue) add(n node) {
	q.q = append(q.q, n)
}

func (q *queue) remove() (node, bool) {
	if len(q.q) == 0 {
		return node{}, false
	}
	n := q.q[0]
	q.q = q.q[1:]
	return n, true
}

func shortestPathBinaryMatrix(grid [][]int) int {
	size := len(grid)
	v := newVisitedList(size)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 0 {
				v.addNew(point{i, j})
			}
		}
	}
	if !v.addNew(point{0, 0}) {
		return -1
	}
	if size == 1 {
		return 1
	}
	q := newQueue()
	q.add(node{0, 0, 0})
	for {
		n, ok := q.remove()
		if !ok {
			return -1
		}
		children := n.children(&v)
		for _, child := range children {
			if child.i == size-1 && child.j == size-1 {
				return child.N + 1
			}
			q.add(child)
		}
	}
}

func main() {
	mat := [][]int{
		/*
			{0, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 0},
			{0, 1, 1, 1, 1, 0},
			{0, 0, 0, 0, 1, 0},
			{0, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 0},
		*/
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	res := shortestPathBinaryMatrix(mat)
	fmt.Println(res)
}
