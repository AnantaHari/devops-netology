package main

import (
	"fmt"
)

func main() {
	fmt.Print("Введите длину в футах: ")
	var input float64
	fmt.Scanf("%f", &input)

	fmt.Println(convert_fm(input))

	// output := input * 0.3048

	// fmt.Printf("%v футов равно %.4v метров.\n", input, output)

	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}

	fmt.Printf("Самое маленькое число - %v\n", little_element(x))

	fmt.Println(div_three())
}

func little_element(x []int) int {
	s := x[0]
	for _, xx := range x {
		if s > xx {
			s = xx
		}
	}
	return s
}

func div_three() []int {
	var numbers []int
	for i := 1; i < 101; i++ {
		if i%3 == 0 {
			numbers = append(numbers, i)
		}
	}
	return numbers
}
