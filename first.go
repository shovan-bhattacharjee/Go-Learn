package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	// variable declaration
	var x int
	x = 5
	fmt.Println("The value of x is:", x)

	// variable declaration with inference
	var y = 20
	fmt.Println("The value of y is:", y)

	// this only works inside functions, not at the package level
	// also declares and initializes the variable in one step
	a := 10 // shorthand variable declaration and inference

	fmt.Println("The value of a is:", a)
	fmt.Printf("The value of a is: %d\n", a)

	// multiple variable declaration
	var b, c bool = true, false
	fmt.Println("The value of b is:", b)
	fmt.Println("The value of c is:", c)

	// multiple variable declaration with inference
	d, e := "Hello", "Go"
	fmt.Println("The value of d is:", d)
	fmt.Println("The value of e is:", e)

	var (
		f int
		g int    = 1
		h string = "hello"
	)

	fmt.Println("The value of f is:", f)
	fmt.Println("The value of g is:", g)
	fmt.Println("The value of h is:", h)

	const pi = 3.14
	// pi = 3.14159 // This will cause a compile-time error because pi is a constant
	fmt.Println("The value of pi is:", pi)
}
