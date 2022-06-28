package shapes

type Shape interface {
	Area() float64
}

type Shapes []Shape

func (s Shapes) LargestArea() float64 {
	temp := s[0].Area()
	for _, shape := range s {
		if shape.Area() > temp {
			temp = shape.Area()
		}
	}
	return temp
}
