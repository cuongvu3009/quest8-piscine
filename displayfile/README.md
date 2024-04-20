## Instructions

### Allowed functions: fmt.*, os.*, io/ioutil.*

Write a program that displays the content of a file given as an argument on the standard output.

## Usage

```sh
$ go run .
File name missing
$ echo "Almost there!!" > quest8.txt
$ go run . quest8.txt main.go
Too many arguments
$ go run . quest8.txt
Almost there!!
```