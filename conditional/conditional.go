package main

import "fmt"

func main() {
	age := 14

	if age < 18 {
		fmt.Println("You are not eligible to vote.")
	} else if age >= 18 && age < 65 {
		fmt.Println("You are eligible to vote.")
	} else {
		fmt.Println("You are a senior and eligible to vote.")
	}

	// switch statement
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Today is Monday.")
	case "Tuesday":
		fmt.Println("Today is Tuesday.")
	default:
		fmt.Println("Today is not Monday or Tuesday.")
	}

	// switch with multiple cases
	grade := "A"
	switch grade {
	case "A", "B", "C":
		fmt.Println("You passed.")
	case "D", "F":
		fmt.Println("You failed.")
	default:
		fmt.Println("Invalid grade.")
	}

	// switch without a condition
	switch {
	case age < 18:
		fmt.Println("You are not eligible to vote.")
	case age >= 18 && age < 65:
		fmt.Println("You are eligible to vote.")
	default:
		fmt.Println("You are a senior and eligible to vote.")
	}

	// switch with fallthrough
	switch {
	case age < 18:
		fmt.Println("You are not eligible to vote.")
		fallthrough
	case age >= 18 && age < 65:
		fmt.Println("You are eligible to vote.")
		fallthrough
	default:
		fmt.Println("You are a senior and eligible to vote.")
	}
}