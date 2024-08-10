package main

import (
	"fmt"
	"github.com/Laky-64/http"
)

func main() {
	res, err := http.ExecuteRequest(
		"https://httpbin.org/cookies",
		http.Method("GET"),
		http.Cookies(
			map[string]string{
				"Jar": "Biscuit",
			},
		),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
