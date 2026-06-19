## Week 1: Go Fundamentals

### Concepts Learned

- Concept 1: Types, Variables, Control Flow
- Concept 2: String Manipulation
- Concept 3: if, else-if, switch case (with fallthrough), defer

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

### Challenges & Solutions

- Challenge: [Describe a problem you encountered]
- Solution: [How you solved it]

### Resources Used

- [Link to helpful resource]
- [Another helpful resource]

### Next Week's Focus

- What to learn next
- Project next steps
