package library

import "fmt"

type Student struct {
    Name  string
    Grade int
}

var Teacher =  struct {
    Name  string
}{}

func SayHello(name string) {
	fmt.Println("hello")
	introduce(name)
}

func introduce(name string) {
    fmt.Println("nama saya", name)
}

func init() {
    Teacher.Name = "Guru Ageung"

	fmt.Println("--> library/library.go imported")
	fmt.Println(Teacher.Name)
}