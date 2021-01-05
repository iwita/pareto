package pareto

type ParetoFrontierGenerator struct {
	points []Point
}

type Point struct {
	isFront bool
	x1      float64
	x2      float64
	Id      int
}
