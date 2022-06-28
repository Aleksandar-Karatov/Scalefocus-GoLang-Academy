package square

type Square struct {
	SideValue float64
}

func (r *Square) Area() float64 {
	return r.SideValue * r.SideValue
}
