package main

import (
	. "akses/library" // . untuk akses SayHello tanpa awalan library.SayHello
	f "fmt" // alias menjadi f
)

func main() {
	SayHello("cipta")
	//library.introduce("cipta") jika function diawali huruf kecil itu adalah private akses tidak dapat dipanggil disini
	var s1 = Student{"cipta ageung", 21}
	f.Println("name ", s1.Name)
	f.Println("grade", s1.Grade)

	haloTeman("soni")
}