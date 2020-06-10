package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(p.Distance(q))

	names := StringSlice{"122143", "dffdgfdg", "asdfsdgfdsg"}
	sort.Sort(names)
	fmt.Println(names)

}

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
