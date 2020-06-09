package something

import "fmt"

// lowercase : private
func sayBye() {
	fmt.Println("Bye")
}

// Uppercase : export func
func SayHello() {
	fmt.Println("Hello")
}
