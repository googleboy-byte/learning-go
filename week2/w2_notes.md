## Week 1: Go Fundamentals

### Concepts Learned

- Concept 1: Functions

### Code Snippet of the Week

```go
// function call by reference

func multiplyByPointer(a, b *int) int {
	(*a)++
	(*b)++
	return (*a) * (*b)
} // changes inside the function reflect in the caller's variables

...

fmt.Println("Calling multiplyByPointer function")
a := 1
b := 2
fmt.Println(multiplyByPointer(&a, &b))

```

### Challenges & Solutions

### Resources Used

- [go functions](https://geeksforgeekgs.org/go-language/functions-in-go-language/)

### Next Week's Focus

- What to learn next
- Project next steps
