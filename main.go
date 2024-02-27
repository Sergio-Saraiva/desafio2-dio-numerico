package main

import "fmt"

func main() {
	divisivelPor3()
	pin()
}

func divisivelPor3() {
	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func pin() {
	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println(i, "Pin")
		}

		if i%5 == 0 {
			fmt.Println(i, "Pan")
		}
	}
}
