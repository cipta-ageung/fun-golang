package main

import "fmt"
import "reflect"

func main() {
    var number = 23
    var reflectValue = reflect.ValueOf(number)

    fmt.Println("tipe  variabel :", reflectValue.Type())

    if reflectValue.Kind() == reflect.Int {
        fmt.Println("nilai variabel :", reflectValue.Int())
	}
	
	var s1 = &student{Name: "cipta", Grade: 2}
	s1.getPropertyInfo()
	
	var s2 = &student{Name: "cipta ageung", Grade: 2}
    fmt.Println("nama :", s2.Name)

    var reflectValue2 = reflect.ValueOf(s2)
    var method = reflectValue2.MethodByName("SetName")
    method.Call([]reflect.Value{
        reflect.ValueOf("cipta"),
    })

    fmt.Println("nama :", s2.Name)
}