package main

import "fmt"

func main() {
	// fmt.Println("Hello, World!")

	const kuendigung = "sofort"

	var ages [3]int = [3]int{20, 25, 30}
	noombers := [5]float64{1.2, 1.3, 1.4, 1.6}

	fmt.Println(ages[0])
	fmt.Println(kuendigung)
	fmt.Println(noombers)

	noombers[4] = 8.8
	fmt.Println(noombers)

	// slices
	var testslice = []int{10, 5, 2}
	testslice = append(testslice, 99)

	fmt.Println(testslice, len(testslice))
}
