package main

import "fmt"

func main() {
	// vectores son listas de elementos iguales
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println("Numbers:", numbers)
	fmt.Println("First number:", numbers[0])
	fmt.Println("Second number:", numbers[1])
	fmt.Println("Last number:", numbers[len(numbers)-1])

	var slice1 []int
	fmt.Println("Empty slice:", slice1)
	slice1 = append(slice1, 12, 22, 33)
	fmt.Println("Updated slice:", slice1)

	slice2 := slice1[1 : len(slice1)-1]
	fmt.Println("New slice:", slice2)
	// inicio:final
	fmt.Println("Full slice:", slice1[:])

	slice3 := make([]int, 3)
	fmt.Println("New slice with make:", slice3)
}
