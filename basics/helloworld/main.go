package main

import (
	"fmt"
	"strings"

	"github.com/harryjoo/learngo/basics/helloworld/something"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func lenAndUpper2(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done") // defer : run defer code, after this function is executed...
	length = len(name)
	uppercase = strings.ToUpper(name)
	return // naked return
}

func main() {
	fmt.Println("Hello world!")
	something.SayHello()

	//const name string = "nico" (Constant)
	//var name string = "nico" (Variable)

	// var name string = "nico" 의 축약형
	// Go가 name의 type을 찾아 정해준다
	// 축약형은 function 안에서만 가능함
	name := "nico"
	name = "lynn"

	fmt.Println(name)

	fmt.Println(multiply(2, 2))

	totalLength, upperName := lenAndUpper("nico")
	fmt.Println(totalLength, upperName)

	// lenAndUpper은 2개의 value를 리턴함
	// _ : ignore value
	totalLength2, _ := lenAndUpper("nico")
	fmt.Println(totalLength2)

	repeatMe("nico", "lynn", "dal", "marl", "flynn")

	totalLength3, uppserName3 := lenAndUpper2("nico")
	fmt.Println(totalLength3, uppserName3)
}
