package main

import (
    "fmt"
    "math/rand"
	"strings"
	"math"
)

func main(){
	var names = []string{"Cipta", "Ageung"}
	printMessage("halo", names)
	
	var randomValue int
	randomValue = randomWithRange(2, 10)
	fmt.Println("anda mendapatkan nomer antri : ", randomValue)
	
	var diameter float64 = 7
    var area, circumference = calculate(diameter)

    fmt.Printf("Anda mendapatkan dorprize bola dengan luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("Anda mendapatkan dorprize bola dengan keliling lingkaran\t: %.2f \n", circumference)
	
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	//var avg = calculate2(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var avg = calculate2(numbers...)
    var msg = fmt.Sprintf("Rata-rata total antrian : %.2f", avg)
	fmt.Println(msg)
	
	// kasus lain untuk contoh function
	var getMinMax = func(n []int) (int, int) {
        var min, max int
        for i, e := range n {
            switch {
            case i == 0:
                max, min = e, e
            case e > max:
                max = e
            case e < min:
                min = e
            }
        }
        return min, max
    }

    var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)
	
	var data = []string{"cipta", "ageung", "mahdiar"}
    var dataContainsO = filter(data, func(each string) bool {
        return strings.Contains(each, "g")
    })
    var dataLenght5 = filter(data, func(each string) bool {
        return len(each) == 7
    })

    fmt.Println("Data diri \t\t:", data)
    // Data diri : [cipta ageung mahdiar]

    fmt.Println("filter ada huruf \"a\"\t:", dataContainsO)
    // filter ada huruf "o" : [ageung]

    fmt.Println("filter jumlah huruf \"7\"\t:", dataLenght5)
    // filter jumlah huruf "5" : [ageung mahdiar]
}

func printMessage(message string, names []string) {
	if message == "" {
        fmt.Println("Tidak ada pesan")
        return
    }
    var nameString = strings.Join(names, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
    var value = rand.Int() % (max - min + 1) + min
    return value
}

func calculate(d float64) (float64, float64) {
    // hitung luas
    var area = math.Pi * math.Pow(d / 2, 2)
    // hitung keliling
    var circumference = math.Pi * d

    // kembalikan 2 nilai
    return area, circumference
}

func calculate2(numbers ...int) float64 {
    var total int = 0
    for _, number := range numbers {
        total += number
    }

    var avg = float64(total) / float64(len(numbers))
    return avg
}

func filter(data []string, callback func(string) bool) []string {
    var result []string
    for _, each := range data {
        if filtered := callback(each); filtered {
            result = append(result, each)
        }
    }
    return result
}