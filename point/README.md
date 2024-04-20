## Instructions

### Allowed functions: github.com/01-edu/z01.PrintRune

1. Create a new directory called `point`.

2. Copy the code below into a file called `main.go` inside the `point` directory.

3. Apply the necessary changes so that the program works.

4. Ensure that the function `setPoint()` works with `int`.


## Code to be copied

```go
func setPoint(ptr *point) {
    ptr.x = 42
    ptr.y = 21
}

func main() {
    points := &point{}

    setPoint(points)

    fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
```

```go
$ go run .
x = 42, y = 21
```