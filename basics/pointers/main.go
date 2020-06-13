package main

import "fmt"

func main() {
	// Go lang : low-level programming
	a := 2
	b := &a // a의 주소값을 저장

	fmt.Println(a)
	fmt.Println(&a, b)
	fmt.Println(*b)

	a = 5
	fmt.Println(*b)

	*b = 20 // b는 a를 살펴보는 pointer
	fmt.Println(a)
	fmt.Println(*b)
}
