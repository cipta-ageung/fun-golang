package main

import "fmt"
import "strings"

func main() {
    var bangunRuang hitung = &kubus{4}

    fmt.Println("===== kubus")
    fmt.Println("luas      :", bangunRuang.luas())
    fmt.Println("keliling  :", bangunRuang.keliling())
	fmt.Println("volume    :", bangunRuang.volume())
	
	var secret interface{}

    secret = 2
    var number = secret.(int) * 10
    fmt.Println(secret, "multiplied by 10 is :", number)

    secret = []string{"apple", "manggo", "banana"}
    var gruits = strings.Join(secret.([]string), ", ")
	fmt.Println(gruits, "is my favorite fruits")
	
	var person1 interface{} = &person{name: "wick", age: 27}
	var name = person1.(*person).name
	fmt.Println(name)


	var person2 = []map[string]interface{}{
		{"name": "cipta ageung mahdiar", "age": 29},
		{"name": "soni setiabudi", "age": 28},
		{"name": "sabil hasni", "age": 29},
	}
	
	for _, each := range person2 {
		fmt.Println(each["name"], "age is", each["age"])
	}

	var fruits = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}
	
	for _, each := range fruits {
		fmt.Println(each)
	}
}

type person struct {
    name string
    age  int
}