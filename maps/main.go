package main

import "fmt"

func main() {
	// los maps son diccionarios clave valor
	m1 := make(map[int]string)
	m1[1] = "Golang"
	m1[2] = "Java"
	m1[3] = "Python"

	fmt.Println("Map m1:", m1)
	fmt.Println("Value for key 1:", m1[1])
	fmt.Println("Value for key 2:", m1[2])
	fmt.Println("Value for key 3:", m1[3])

	m2 := map[string]int{
		"one": 1,
		"two": 2,
		"ten": 10,
	}
	fmt.Println("Map m2:", m2)
	delete(m2, "two")
	fmt.Println("Map m2 after deletion:", m2)

	value, exists := m2["ten"]
	fmt.Println("Value for key 'ten':", value, "Exists:", exists)

	value, exists = m2["two"]
	fmt.Println("Value for key 'two':", value, "Exists:", exists)

	m3 := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	fmt.Println("Map m3:", m3)
}
