package main

import "fmt"

func main(){
	var saudara map[string]int
	saudara = map[string]int{}

	saudara["adik"] = 2
	saudara["kakak"] = 1

	fmt.Println("adik\t:", saudara["adik"]) // adik 2
	fmt.Println("kakak\t:",     saudara["kakak"])     // kakak 1

	var teman = map[string]int{
		"dekat":  5,
		"jauh": 15,
		"mesra":    2,
		"pacar":    1,
	}
	
	for key, val := range teman {
		fmt.Println(key, "\t:", val)
	}

	fmt.Print(len(teman),"\t:\t") // 4
	fmt.Println(teman)

	delete(teman, "pacar")

	fmt.Print(len(teman),"\t:\t") // 3
	fmt.Println(teman)

	var value, isExist = teman["pacar"]

	if isExist {
	    fmt.Println(value)
	} else {
	    fmt.Println("sudah tidak punya pacar / putus")
	}

	var mantans = []map[string]string{
		{"nama": "silvia",   "alamat": "indonesia", "mantanKe": "1"},
		{"nama": "claudia",   "alamat": "amerika", "mantanKe": "10"},
	}

	fmt.Println(mantans)
	for _, mantan := range mantans {
		fmt.Println(mantan["nama"], mantan["mantanKe"])
	}
}