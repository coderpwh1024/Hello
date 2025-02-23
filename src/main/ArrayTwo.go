package main

import "fmt"

func main() {

	var i, j, k int

	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] =%f\n", i, balance[i])
	}

	fmt.Println("------------------------------------")

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	for j = 0; j < 5; j++ {
		fmt.Printf("balance2[%d] =%f\n", j, balance2[j])
	}

	fmt.Println("------------------------------------")

	balance3 := [5]float32{1: 2.0, 3: 7.0}

	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d]=%f\n", k, balance3[k])
	}

}
