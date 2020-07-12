package main

import "fmt"

//배열
//값의 크기가 정해지지 않은 것
//append 하는 것
func main() {
	names := []string{"nico", "lynn", "dal"}
	names = append(names, "dsadwgwzcsd")
	fmt.Println(names)
}
