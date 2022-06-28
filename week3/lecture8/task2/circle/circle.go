package circle

type Circle struct {
	Radius float64
}

func (r *Circle) Area() float64 {
	return 3.14 * r.Radius * r.Radius
}
