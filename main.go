package main

import (
	"fmt"

	"github.com/iwita/pareto/pareto"
)

// func NewSolver(items ...Objectives) *ParetoSolver {
// 	if len(items) <= 1 {
// 		fmt.Println("Not a multi-objective problem")
// 	}
// 	for _, item := range items {
// 		item.paretoOptimal = true
// 	}

// 	return &ParetoSolver{objectives: items}
// }

// func (p *ParetoSolver) start() {

// }

// ebox = math.floor(objectives[x]/epsilons[x])

// func Transpose(a [][]bool) {
// 	n := len(a)
// 	b := make([][]bool, n)
// 	for i := 0; i < n; i++ {
// 		b[i] = make([]bool, n)
// 		for j := 0; j < n; j++ {
// 			b[i][j] = a[j][i]
// 		}
// 	}
// 	copy(a, b)
// }

func main() {
	// input := []Point{
	// 	Point{x1: 1, x2: 8},
	// 	Point{x1: 2, x2: 7},
	// 	Point{x1: 2.9, x2: 100},
	// 	Point{x1: 3, x2: 11},
	// 	Point{x1: 4, x2: 2},
	// 	Point{x1: 5, x2: 20},
	// }

	attr1 := []float64{1, 2, 3, 4, 5}
	attr2 := []float64{8, 7, 11, 2, 20}

	p, _ := pareto.New(attr1, attr2)
	fmt.Println(p)
}
