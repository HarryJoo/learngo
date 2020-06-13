package main

import "fmt"

func canIDrink2(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func canIDrink3(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
}

func canIDrink4(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDrink2(16))
	fmt.Println(canIDrink3(16))
	fmt.Println(canIDrink4(16))
}
