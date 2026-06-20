package main

import (
	"fmt"
	"os"

	"github.com/Knetic/govaluate"
)

func main() {

	cli_expression := os.Args[1]

	expression, err := govaluate.NewEvaluableExpression(cli_expression)
	if err != nil {
		fmt.Println("Error: Invalid expression", err)
		return
	}

	result, err := expression.Evaluate(nil)
	// if the expression contains variables,
	// then the evaluate function expects a map of strings to interfaces as a parameter.
	// in this case, there are no variables, so nil is passed as a parameter.
	if err != nil {
		fmt.Println("Error: Invalid expression", err)
		return
	}

	fmt.Println("The result of", cli_expression, "is", result)
}

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {

// 	// calculator with command line arguments

// 	c_arg1, err := strconv.Atoi(os.Args[1]) // convert to type int
// 	if err != nil {
// 		fmt.Println("Error: Invalid number", err)
// 		return
// 	}

// 	c_arg2 := os.Args[2]

// 	c_arg3, err := strconv.Atoi(os.Args[3])
// 	if err != nil {
// 		fmt.Println("Error: Invalid number", err)
// 		return
// 	}

// 	switch c_arg2 {
// 	case "+":
// 		fmt.Println(c_arg1 + c_arg3)
// 	case "-":
// 		fmt.Println(c_arg1 - c_arg3)
// 	case "*":
// 		fmt.Println(c_arg1 * c_arg3)
// 	case "/":
// 		fmt.Println(c_arg1 / c_arg3)
// 	default:
// 		fmt.Println("Invalid operator")
// 	}
// }
