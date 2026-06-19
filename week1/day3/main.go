package main

import "fmt"

func main() {
	// long declaration
	var x int = 5
	var y string = "hello"
	var y1 = "hello"
	var z bool = true

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(y1)
	fmt.Println(z)

	// short declaration
	w := 10
	fmt.Println(w)

	// multiple short declarations
	a, b := 10, 20
	fmt.Println(a, b)

	a = 30
	fmt.Println(a, b)

	const pi = 3.14159
	fmt.Println(pi)

	age := 24

	if age > 21 {
		fmt.Println("Drinking age")
	} else {
		fmt.Println("Not drinking age")
	}

	// go requires explicit type conversion, unlike python
	v1 := 10
	v2 := 10.5
	fmt.Println(float64(v1) + v2)

	// any new variable declaration requires := in short form
	sum := float64(v1) + v2
	fmt.Println(sum)

	// new concept. cool. iota. increases value by 1. indexing starts at 0
	// by default, iota increases by 1. so multiplying by 2 increases by 2
	// multiplying iota by any 'n' increases counter by n
	// adding any 'x' to iota increases counter by x. can set starting index
	// using this added offset.

	const (
		Monday = (iota * 2) + 10
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	fmt.Println("Monday", Monday)
	fmt.Println("Tuesday", Tuesday)
	fmt.Println("Wednesday", Wednesday)

	// zero value of variable types

	var integer int
	var str string
	var flt float64
	var bln bool

	fmt.Println(integer, str, flt, bln)

	// string manipulation

	// concatenation
	str1 := "Hello"
	str2 := "World"

	concat := str1 + ", " + str2
	fmt.Println(concat)

	// length of string
	fmt.Println(len(concat))

	// string indexing
	fmt.Println(concat[0])

	// string formatting with Sprintf.
	name := "Mainak"
	name1 := "Surname"
	age = 20
	class := 12

	msg := fmt.Sprintf(
		"My name is %s %s and I am %d. I read in class %d",
		name,
		name1,
		age,
		class,
	)

	fmt.Println(msg)

}
