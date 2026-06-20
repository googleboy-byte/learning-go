## Week 1: Go Fundamentals

### Concepts Learned

- Concept 1: Types, Variables, Control Flow
- Concept 2: String Manipulation
- Concept 3: if, else-if, switch case (with fallthrough), defer
- Concept 4: Command Line Arguments, type conversion (str to int)
- Concept 5: package, import, go get

### Code Snippet of the Week

```go
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
```

```go
// deferred execution
	defer fmt.Println("hello") // this will execute at the end of the function
	fmt.Println("world")
```

```go
cli_expression := os.Args[1]

expression, err := govaluate.NewEvaluableExpression(cli_expression)
if err != nil {
    fmt.Println("Error: Invalid expression", err)
    return
}

result, err := expression.Evaluate(nil)
```

### Challenges & Solutions

- Challenge: [string to int conversion for cli args]
- Solution: [strconv.Atoi]

- Challenge: [evaluate formula from cli args]
- Solution: [govaluate]

### Resources Used

- [string to int conversion for cli args](https://stackoverflow.com/questions/4278430/convert-string-to-integer-type-in-go)
- [how to evaluate a formula in go](https://stackoverflow.com/questions/23923383/evaluate-formula-in-go)

### Next Week's Focus

- What to learn next
- Project next steps
