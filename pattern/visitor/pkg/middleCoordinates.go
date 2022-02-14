package pkg

import (
	"fmt"
)

type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) VisitForSquare(s *Square) {
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) VisitForCircle(c *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}

func (a *MiddleCoordinates) VisitForRectangle(r *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}
