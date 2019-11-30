package main

import "fmt"
import "os"

// run dengan go run bab45.go banana potato "ice cream"

func main() {
	var argsRaw = os.Args
	fmt.Printf("-> %#v\n", argsRaw)
	// []string{".../bab45", "banana", "potato", "ice cream"}

	var args = argsRaw[1:]
	fmt.Printf("-> %#v\n", args)
	// []string{"banana", "potato"}
}