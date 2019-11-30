package main

import "fmt"

func main(){
	var left = true
	var right = false

	var leftAndRight = left && right
	fmt.Printf("left & right : \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right: \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left: \t\t(%t) \n", leftReverse)
}