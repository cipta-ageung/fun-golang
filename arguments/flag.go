package main

import "flag"
import "fmt"

// run dengan go run bab45.go -name="john wick" -age=28
// atau build dengan go build flag.go && ./flag.go -name="john wick" -age=28 or ./flag.go --help

func main() {
	var name = flag.String("name", "anonymous", "type your name")
	var age = flag.Int64("age", 25, "type your age")

	flag.Parse()
	fmt.Printf("name\t: %s\n", *name)
	fmt.Printf("age\t: %d\n", *age)
}