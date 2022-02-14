package pkg

import (
	"fmt"
)

type AreaCalculation struct {
	area int
}

func (a *AreaCalculation) VisitForSquare(s *Square) {
	fmt.Println("Calculating area for square")
}

func (a *AreaCalculation) VisitForCircle(c *Circle) {
	fmt.Println("Calculating area for circle")
}

func (a *AreaCalculation) VisitForRectangle(r *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}
