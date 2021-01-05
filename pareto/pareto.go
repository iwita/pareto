package pareto

import (
	"sort"
)

func New(input map[string][2]float64) (*ParetoFrontierGenerator, error) {

	// if len(attr1) != len(attr2) {
	// 	return nil, fmt.Errorf("Wrong input")
	// }
	var i int
	points := make([]Point, len(input))
	for k, v := range input {
		points[i].x1 = v[0]
		points[i].x2 = v[1]
		points[i].isFront = true
		points[i].Id = k
		i++
	}

	return &ParetoFrontierGenerator{points: points}, nil
}

func (p *ParetoFrontierGenerator) Exec() []Point {
	return getFront(p.points)
}

func getFront(points []Point) []Point {
	sort.Slice(points, func(i, j int) bool {
		return points[i].x1 < points[j].x1
	})
	//fmt.Println(points)
	ret := make([]Point, 0)
	for idx1 := 1; idx1 < len(points); idx1++ {
		for idx2 := 0; idx2 < idx1; idx2++ {
			if points[idx2].x2 < points[idx1].x2 {
				points[idx1].isFront = false
				break
			}
		}
	}
	for _, p := range points {
		if p.isFront == true {
			ret = append(ret, p)
		}
	}
	return ret
}
