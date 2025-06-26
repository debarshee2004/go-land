package main

import "fmt"

// A struct type for demonstration
type Rectangle struct {
	Width  float64
	Height float64
}

// ----------------------------
// Method with value receiver
func (r Rectangle) Area() float64 {
	// This method does not modify the original struct
	return r.Width * r.Height
}

// ----------------------------
// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
	// This method modifies the actual struct using pointer receiver
	r.Width *= factor
	r.Height *= factor
}

// Another struct to show method sets
type Circle struct {
	Radius float64
}

// Method on Circle
func (c Circle) Circumference() float64 {
	return 2 * 3.1415 * c.Radius
}

// ----------------------------
// Method that returns formatted string
func (r Rectangle) Description() string {
	return fmt.Sprintf("Rectangle [%.2f x %.2f]", r.Width, r.Height)
}

func main() {
	fmt.Println("Welcome to the Methods section of Go programming!")

	// ----------------------------
	// Example 1: Method with value receiver
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("\nExample 1: Area method with value receiver")
	area := rect.Area()
	fmt.Printf("Area of rectangle: %.2f\n", area)

	// ----------------------------
	// Example 2: Method with pointer receiver (modifies struct)
	fmt.Println("\nExample 2: Scale method with pointer receiver")
	rect.Scale(2)
	fmt.Printf("Scaled rectangle: %+v\n", rect)
	fmt.Printf("New area after scaling: %.2f\n", rect.Area())

	// ----------------------------
	// Example 3: Another type with method
	fmt.Println("\nExample 3: Method on Circle struct")
	circle := Circle{Radius: 7}
	fmt.Printf("Circumference of circle: %.2f\n", circle.Circumference())

	// ----------------------------
	// Example 4: Method returning a string
	fmt.Println("\nExample 4: Method that returns a description")
	fmt.Println(rect.Description())
}
