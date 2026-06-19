package main

import (
	"fmt"
) // here is how to do multiple imports

func main() {

	// if statement:
	a := 1
	if a > 0 {
		fmt.Println("a is greater than 0")
	} else if a == 0 {
		fmt.Println("a is equal to 0")
	} else {
		fmt.Println("a is less than 0")
	}

	// go allows initialization in the if condition:
	if x := 5; x > 3 {
		fmt.Println("x is greater than 3")
	}

	// for loop: type 1
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for loop: type 2
	x := 10
	for x > 0 {
		fmt.Println(x)
		x--
	}

	// for loop: type 3
	for {
		fmt.Println("infinite loop")
		break
	}

	// switch statements with fallthrough differences
	casenum := 3
	switch {
	case casenum == 1:
		fmt.Println("one")
		fallthrough // like python's continue. but in reverse. executes next case too
	case casenum == 2:
		fmt.Println("two")
		fallthrough
	case casenum == 3:
		fmt.Println("three")
		fallthrough
	default:
		fmt.Println("default")
	}

	// deferred execution
	defer fmt.Println("hello") // this will execute at the end of the function
	fmt.Println("world")

	// another example
	fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("three")

}
