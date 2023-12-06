package main

import (
	"container/heap"
	_ "embed"
	"strings"
)

//go:embed "input.txt"
var input string

// An Item is something we manage in a priority queue.
type Item struct {
	Coord    [2]int // The value of the item; arbitrary.
	F        int
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Estimation(p [2]int, q [2]int, ch byte) int {
	diff := abs(p[0]-q[0]) + abs(p[1]-q[1])
	chdiff := 'z' - ch
	return max(diff, int(chdiff))
}

func read(s string) ([2]int, [2]int, [][]byte) {
	start := [2]int{}
	end := [2]int{}
	lines := strings.Split(s, "\n")
	// sz := len(lines)
	arr := make([][]byte, 0)
	row := 0
	for k, l := range lines {
		col := 0
		arr = append(arr, []byte{})

		for i := 0; i < len(l); i++ {
			ch := l[i]
			if ch == 'S' {
				ch = 'a'
				start[1] = col
				start[0] = row
			} else if ch == 'E' {
				ch = 'z'
				end[1] = col
				end[0] = row
			}

			arr[k] = append(arr[k], ch)
			col++
		}
		row++
	}
	// fmt.Printf("%+v %v %v", arr, start, end)
	return start, end, arr
}

func shortest(start [2]int, end [2]int, arr [][]byte) int {

	close := map[[2]int]struct{}{}
	open := make(PriorityQueue, 1)
	startEstimation := Estimation(start, end, 'a')
	open[0] = &Item{Coord: start, F: 0, priority: -startEstimation, index: 0}
	heap.Init(&open)

	ydim := len(arr)
	xdim := len(arr[0])
	for {
		if len( open) == 0 {
			return 100000000;
		}
		el := heap.Pop(&open).(*Item)
		if _, found := close[el.Coord]; found {
			continue
		}
		if el.Coord[0] == end[0] && el.Coord[1] == end[1] {
			return el.F
		}

		for i := 0; i < 4; i++ {
			dx := 0
			dy := 0
			switch i {
			case 0:
				dx = -1
			case 1:
				dx = 1
			case 2:
				dy = -1
			case 3:
				dy = 1

			}

			y := el.Coord[0] + dy
			x := el.Coord[1] + dx
			if x < 0 || x >= xdim {
				continue
			}
			if y < 0 || y >= ydim {
				continue
			}

			c := [2]int{y, x}

			if _, found := close[c]; found {
				continue
			}
			oldch := int(arr[el.Coord[0]][el.Coord[1]])
			newch := int(arr[y][x])
			if newch-oldch > 1 {
				continue
			}

			g := Estimation(c, end, arr[y][x])

			heap.Push(&open, &Item{Coord: c, F: el.F + 1, priority: -el.F - 1 - g})
		}
		close[el.Coord] = struct{}{}
	}
}

func main() {
	start, end, arr := read(input)
	println(shortest(start, end, arr))
	ydim := len(arr)
	xdim := len(arr[0])
	s := xdim * ydim
	for y := 0; y < ydim; y++ {
		for x:= 0; x < xdim; x++ {
			if arr[y][x] == 'a' {
				s = min(shortest( [2]int{y,x}, end, arr),s)
			}
		}
	}
	println(s)
}
