package main

import "fmt"

func main(){
	for i:=0;i<5;i++{
		if i % 2 == 1 {
			continue
		}

		if i > 3 {
			break
		}

		fmt.Println("\nPerulangan ke-",i)
	}
}