package main

import "fmt"

//map

//for문과 range 사용
func main() {
	nico := map[string]string{"name": "nico", "age": "12"}

	for key, value := range nico {
		fmt.Println(key, value)
	}

}
