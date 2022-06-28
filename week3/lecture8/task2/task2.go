package main

import (
	"fmt"
	"week3Lecture8Task2/week3/lecture8/task2/circle"
	"week3Lecture8Task2/week3/lecture8/task2/shapes"
	"week3Lecture8Task2/week3/lecture8/task2/square"
)

func main() {
	var sq1 square.Square
	sq1.SideValue = 19
	var sq2 square.Square
	sq2.SideValue = 20
	var cir1 circle.Circle
	cir1.Radius = 10
	var cir2 circle.Circle
	cir2.Radius = 15
	var allShapes shapes.Shapes = shapes.Shapes{&sq1, &sq2, &cir1, &cir2}

	fmt.Println("Max area is: ", allShapes.LargestArea())

}
