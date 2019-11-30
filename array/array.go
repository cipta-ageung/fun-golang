package main

import "fmt"

func main(){
	var kelas [6]int
	kelas[0] = 1
	kelas[1] = 2
	kelas[2] = 3
	kelas[3] = 4
	kelas[4] = 5
	kelas[5] = 6
	// kamu dapat initiate dengan cara : 
	// kelas = [6]int{1,2,3,4,5,6}

	for i:=0;i<len(kelas);i++{
		fmt.Println(i, kelas[i])
	}

	var kelas2 = [...]int{1,2,3}

	for i:=0;i<len(kelas2);i++{
		fmt.Println(i, kelas2[i])
	}

	var kelas3 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	for i:=0;i<len(kelas3);i++{
		fmt.Println(i, kelas3[i])
		for j:=0;j<len(kelas3[i]);j++{
			fmt.Println(kelas3[i][j])
		}
	}
	
}