package main

import "fmt"

func add (num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is:", sum)
}

func multiply (num1 int, num2 int) int {
	return num1 * num2
}

func divide (num1 int, num2 int) (int, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return num1 / num2, nil
}

// when func returns multiple values we use () to group them together
func swap (a string, b string) (string, string) {
	return b, a
}

func main() {
	fmt.Println("Hello, Function World!")
	a := 5
	b := 10
	add(a, b)

	result := multiply(a, b)
	fmt.Println("The product of", a, "and", b, "is:", result)

	quotient, err := divide(a, b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("The quotient of", a, "and", b, "is:", quotient)
	}

	c := "Hello"
	d := "World"
	e, f := swap(c, d)
	fmt.Println("After swapping:", e, f)
}