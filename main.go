package main

import "fmt"

//Structs
//데이터를 유연하게 사용하기
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"asdsdfq", "ijoefiunjcsa"}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico)
	fmt.Println(nico.name)

}
