package main

import (
	"fmt"
	"time"
)

func main() {
	// 20초 수행
	//sexyCount("nico")
	//sexyCount("flynn")

	// go routines : 10초 수행
	//go sexyCount("nico")
	//sexyCount("flynn")

	// main 함수 바로 종료
	//go sexyCount("nico")
	//go sexyCount("flynn")

	// 5초 후에 main 함수 종료
	//go sexyCount("nico")
	//go sexyCount("flynn")
	//time.Sleep(time.Second * 5)

	c := make(chan bool) // channel
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
		// result := go isSexy(person) // 이렇게는 사용 불가능
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
	//fmt.Println(<-c) // 채널 2개만 출력해야 하는데 3개를 출력하면 : fatal error: all goroutines are asleep - deadlock!
	time.Sleep(time.Second * 10)
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}
