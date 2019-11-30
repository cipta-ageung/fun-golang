package main

import "fmt"

func main(){
	var s1 student
    s1.name = "cipta ageung mahdiar"
    s1.grade = 1
    fmt.Println("name  :", s1.name)
	fmt.Println("grade :", s1.grade)

	var s2 = student{"cipta", 2}
	var s3 = student{name: "ageung"}
	var s4 = student{}
	s4.name = "mahdiar"
	s4.grade = 2
	fmt.Println("student 1 :", s1.name)
	fmt.Println("student 2 :", s2.name)
	fmt.Println("student 3 :", s3.name)
	fmt.Println("student 4 :", s4.name)

	var s5 = student{name: "cipta", grade: 2}
	var s6 *student = &s5
	fmt.Println("student 5, name :", s5.name)
	fmt.Println("student 6, name :", s6.name)
	s5.name = "ageung"
	fmt.Println("student 5, name :", s5.name)
	fmt.Println("student 6, name :", s6.name)

	var s7 = guru{}
    s7.name = "cipta"
    s7.age = 21
    s7.grade = 2
    fmt.Println("name  :", s7.name)
    fmt.Println("age   :", s7.age)
    fmt.Println("age   :", s7.person.age)
	fmt.Println("grade :", s7.grade)
	
	var s8 = orangtua{}
    s8.name = "cipta"
    s8.age = 21        // umur orang tua
    s8.person.age = 22 // umur person
    fmt.Println(s8.name)
    fmt.Println(s8.age)
	fmt.Println(s8.person.age)
	
	var p1 = person{name: "cipta", age: 21}
	var s9 = guru{person: p1, grade: 2}
	fmt.Println("name  :", s9.name)
	fmt.Println("age   :", s9.age)
	fmt.Println("grade :", s9.grade)

	// anonymous struct tanpa pengisian property
	var s10 = struct {
    	person
    	grade int
	}{}
	s10.person = person{"cipta", 24}
    s10.grade = 2

    fmt.Println("name  :", s10.person.name)
    fmt.Println("age   :", s10.person.age)
    fmt.Println("grade :", s10.grade)

	// anonymous struct dengan pengisian property
	var s11 = struct {
    	person
    	grade int
	}{
    	person: person{"ageung", 25},
    	grade:  2,
	}
	fmt.Println("name  :", s11.person.name)
    fmt.Println("age   :", s11.person.age)
	fmt.Println("grade :", s11.grade)
	
	var allStudents = []person{
		{name: "cipta", age: 27},
		{name: "ageung", age: 28},
		{name: "mahdiar", age: 29},
	}
	for _, student := range allStudents {
		fmt.Println(student.name, "berumur", student.age)
	}

	var allStudents1 = []struct {
		person
		grade int
	}{
		{person: person{"cipta", 27}, grade: 2},
		{person: person{"ageung", 28}, grade: 3},
		{person: person{"mahdiar", 29}, grade: 3},
	}
	
	for _, student := range allStudents1 {
		fmt.Println(student)
	}

	var student = struct {
		grade int
	} {
		12,
	}
	fmt.Println(student)

	// var p1 = struct { name string; age int } { age: 22, name: "wick" }
	// var p2 = struct { name string; age int } { "ethan", 23 }
}

type student struct {
    name string
    grade int
}

type person struct {
    name string
    age  int
}

type guru struct {
    grade int
    person
}

type orangtua struct {
    person
    age   int
    grade int
}