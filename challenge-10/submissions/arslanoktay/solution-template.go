// Package challenge10 contains the solution for Challenge 10.
package challenge10

import (
	"fmt"
	"math"
	"sort"
	// Add any necessary imports here
)

// Shape interface defines methods that all shapes must implement
type Shape interface {
	Area() float64
	Perimeter() float64
	fmt.Stringer // Includes String() string method
}

// Rectangle represents a four-sided shape with perpendicular sides
type Rectangle struct {
	Width  float64
	Height float64
}

// NewRectangle creates a new Rectangle with validation
func NewRectangle(width, height float64) (*Rectangle, error) {
	// TODO: Implement validation and construction
	
	if width == 0 || height == 0 {
	    return nil, fmt.Errorf("Invalid parameters, cannot use zero parameter")
	}
	if width < 0 || height < 0 {
	    return nil, fmt.Errorf("Parameters cannot be negative!")
	}
	
	
	newValue := &Rectangle {
	    Width: width,
	    Height: height,
	}
	return newValue, nil
}

// Area calculates the area of the rectangle
func (r *Rectangle) Area() float64 {
	// TODO: Implement area calculation
	return r.Width * r.Height
}

// Perimeter calculates the perimeter of the rectangle
func (r *Rectangle) Perimeter() float64 {
	// TODO: Implement perimeter calculation
	return 2 * (r.Width + r.Height)
}

// String returns a string representation of the rectangle
func (r *Rectangle) String() string {
	// TODO: Implement string representation
	return fmt.Sprintf("Rectangle(width=%.2f, height=%.2f)", r.Width, r.Height)
}

// Circle represents a perfectly round shape
type Circle struct {
	Radius float64
}

// NewCircle creates a new Circle with validation
func NewCircle(radius float64) (*Circle, error) {
	// TODO: Implement validation and construction
	if radius == 0 || radius < 0 {
	    return nil, fmt.Errorf("Wrong parameter")
	}
	
	circle := &Circle{
	    Radius: radius,
	}
	return circle, nil
}

// Area calculates the area of the circle
func (c *Circle) Area() float64 {
	// TODO: Implement area calculation
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculates the circumference of the circle
func (c *Circle) Perimeter() float64 {
	// TODO: Implement perimeter calculation
	return math.Pi * 2 * c.Radius
}

// String returns a string representation of the circle
func (c *Circle) String() string {
	// TODO: Implement string representation
	return fmt.Sprintf("Circle(Radius=%.2f)",c.Radius)
}

// Triangle represents a three-sided polygon
type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

// NewTriangle creates a new Triangle with validation
func NewTriangle(a, b, c float64) (*Triangle, error) {
	// TODO: Implement validation and construction
	if a + b == c || a + b < c {
	    return nil, fmt.Errorf("Wrong parameters")
	}
	if a == 0 || b == 0 || c == 0 {
	    return nil, fmt.Errorf("zero parameter")
	}
	if a <0 || b < 0 || c < 0 {
	    return nil, fmt.Errorf("negative parameter")
	}
	
	triangle := &Triangle{
	    SideA: a,
	    SideB: b,
	    SideC: c,
	}
	return triangle, nil
}

// Area calculates the area of the triangle using Heron's formula
func (t *Triangle) Area() float64 {
	// TODO: Implement area calculation using Heron's formula
	s := t.Perimeter() / 2
	return math.Sqrt(s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC))
}

// Perimeter calculates the perimeter of the triangle
func (t *Triangle) Perimeter() float64 {
	// TODO: Implement perimeter calculation
	return t.SideA + t.SideB + t.SideC
}

// String returns a string representation of the triangle
func (t *Triangle) String() string {
	// TODO: Implement string representation
	 return fmt.Sprintf("Triangle(sides=%.2f, %.2f, %.2f)", t.SideA, t.SideB, t.SideC)
}

// ShapeCalculator provides utility functions for shapes
type ShapeCalculator struct{
    shapes []Shape
}

// NewShapeCalculator creates a new ShapeCalculator
func NewShapeCalculator() *ShapeCalculator {
	// TODO: Implement constructor
	return &ShapeCalculator{shapes: make([]Shape,0)}
}

// PrintProperties prints the properties of a shape
func (sc *ShapeCalculator) PrintProperties(s Shape) {
	// TODO: Implement printing shape properties
	s.String()
}

// TotalArea calculates the sum of areas of all shapes
func (sc *ShapeCalculator) TotalArea(shapes []Shape) float64 {
	// TODO: Implement total area calculation
	total := 0.0
	for _,s := range shapes {
	    total += s.Area()
	}
	return total
}

// LargestShape finds the shape with the largest area
func (sc *ShapeCalculator) LargestShape(shapes []Shape) Shape {
	// TODO: Implement finding largest shape
	indexBiggest := 0
	biggestArea := 0.0
	for i,s := range shapes {
	    shapeArea := s.Area()
	    if shapeArea > biggestArea {
	        biggestArea = shapeArea
	        indexBiggest = i
	    }
	}
	return shapes[indexBiggest]
}

// SortByArea sorts shapes by area in ascending or descending order
func (sc *ShapeCalculator) SortByArea(shapes []Shape, ascending bool) []Shape {
	// TODO: Implement sorting shapes by area
	sort.Slice(shapes, func(i,j int) bool {
	    if ascending {
	        return shapes[i].Area() < shapes[j].Area()
	    } else {
	        return shapes[i].Area() > shapes[j].Area()
	    }
	})
	
	
	return shapes
} 