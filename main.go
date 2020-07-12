package main

import "fmt"

//for 문 사용하기
func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 24, 235, 324, 23452, 2342345, 72427234)

	fmt.Println(result)

}
