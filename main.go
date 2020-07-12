package main

import "fmt"

// if문 사용하기
// if문 안에 변수 선언해서 핸들링 하기
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
