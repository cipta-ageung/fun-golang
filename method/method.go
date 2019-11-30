package main

import "fmt"
import "strings"

type student struct {
    name  string
    grade int
}

func (s student) sayHello() {
    fmt.Println("halo", s.name)
}

func (s student) getNameAt(i int) string {
    return strings.Split(s.name, " ")[i-1]
}

func (s student) changeName1(name string) {
    fmt.Println("---> on changeName1, name changed to", name)
    s.name = name
}

func (s *student) changeName2(name string) {
    fmt.Println("---> on changeName2, name changed to", name)
    s.name = name
}

func main() {
    var s1 = student{"cipta ageung mahdiar", 21}
    s1.sayHello()

    var name = s1.getNameAt(3)
	fmt.Println("nama panggilan :", name)
	
	s1.changeName1("cipta ageung")
    fmt.Println("s1 after changeName1", s1.name)
    // cipta ageung mahdiar

    s1.changeName2("cipta ageung")
    fmt.Println("s1 after changeName2", s1.name)
    // cipta ageung
}