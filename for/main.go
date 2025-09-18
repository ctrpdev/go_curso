package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum++
	}
	fmt.Println(sum)

	for sum < 100 {
		sum++
	}
	fmt.Println(sum)

	for {
		if sum >= 105 {
			break
		}
		sum++
	}
	fmt.Println(sum)

	array := []int{1, 2, 3, 4, 5, 6, 2, 1, 1}

	// usar el array para acceder al elemento en el indice
	for index := range array {
		fmt.Println(index, ") ", array[index])
	}

	// se obtiene el valor directamente en el for junto con su indice
	for index, value := range array {
		fmt.Println(index, ") ", value)
	}

	// si no se usa el indice, se puede omitir con _
	for _, value := range array {
		fmt.Println("value", value)
	}

	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range map1 {
		fmt.Println(key, ":", value)
	}

	map2 := map[string][]int{
		"A": nil,
		"B": {1, 2, 3},
		"C": {4, 5, 6},
	}
	for key, value := range map2 {
		fmt.Println(key, ":", value)
		for _, v := range value {
			fmt.Println("  -", v)
		}
		fmt.Println()
	}
}
