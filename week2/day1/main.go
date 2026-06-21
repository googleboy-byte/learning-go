package main

import "fmt"

// functions

// basic function

func add(a, b int) int {
	return a + b
} // func, func_name(parameter list) return_type { func_body }

// multiple return values

func swap(a, b int) (int, int) {
	return b, a
} // multiple return values - comma separated

// function call by value

func multiply(a, b int) int {
	a++
	b++
	return a * b
}

// function call by reference

func multiplyByPointer(a, b *int) int {
	(*a)++
	(*b)++
	return (*a) * (*b)
} // changes inside the function reflect in the caller's variables

func main() {

	fmt.Println("Calling add function")
	fmt.Println(add(1, 2))

	fmt.Println("Calling swap function")
	fmt.Println(swap(1, 2))

	fmt.Println("Calling multiply function")
	fmt.Println(multiply(1, 2))

	fmt.Println("Calling multiplyByPointer function")
	a := 1
	b := 2
	fmt.Println(multiplyByPointer(&a, &b))

}
