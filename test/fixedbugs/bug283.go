// compile

// https://golang.org/issue/806
// triggered out of registers on 8g

package bug283

type Point struct {
	x int
	y int
}

func dist(p0, p1 Point) float64 {
	return float64((p0.x-p1.x)*(p0.x-p1.x) + (p0.y-p1.y)*(p0.y-p1.y))
}
