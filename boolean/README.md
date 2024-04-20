## Instructions

### Allowed functions: github.com/01-edu/z01.PrintRune, os.Args

1. Create a new directory called `boolean`.

2. Copy the code below into a file called `main.go` inside the `boolean` directory.

3. Apply the necessary changes for the program to work.

## Code to be copied

```go
func printStr(s string) {
    for _, r := range s {
        z01.PrintRune(r)
    }
    z01.PrintRune('\n')
}

func isEven(nbr int) boolean {
    if even(nbr) == 1 {
        return yes
    } else {
        return no
    }
}

func main() {
    if isEven(lengthOfArg) == 1 {
        printStr(EvenMsg)
    } else {
        printStr(OddMsg)
    }
}
```

```go
$ go run . "not" "odd"
I have an even number of arguments
$ go run . "not even"
I have an odd number of arguments
```