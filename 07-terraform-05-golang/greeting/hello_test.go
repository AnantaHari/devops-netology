package main

import "testing"

func TestHello(t *testing.T) {

	emptyResult := hello("")

	if emptyResult != "Hello Dude!" {
		t.Errorf("hello(\"\") failed, expected %v, got %v", "Hello Dude!", emptyResult)
	}

	result := hello("Mike")

	if result != "Hello Mike!" {
		t.Errorf("hello(\"Mike\") failed, expected %v, got %v", "Hello Dude!", result)
	}

	res := convertfm(3)

	if res != "3 футов равно 0.9145 метров." {
		t.Error("Expected 0.9145, got ", res)
	}
	// x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	//
	// fmt.Printf("Самое маленькое число - %v\n", little_element(x))
	//
	// fmt.Println(div_three())
}

// func little_element(x []int) int {
// 	s := x[0]
// 	for _, xx := range x {
// 		if s > xx {
// 			s = xx
// 		}
// 	}
// 	return s
// }
//
// func div_three() []int {
// 	var numbers []int
// 	for i := 1; i < 101; i++ {
// 		if i%3 == 0 {
// 			numbers = append(numbers, i)
// 		}
// 	}
// 	return numbers
// }
