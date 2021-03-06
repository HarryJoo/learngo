package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	//var results = map[string]string{}
	var results = make(map[string]string)

	// 참고) 초기화 되지 않은 map에는 값을 넣을 수 없음
	//var results map[string]string // panic: assignment to entry in nil map

	urls := []string{
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
