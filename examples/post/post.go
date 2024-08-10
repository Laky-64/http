package main

import (
	"fmt"
	"github.com/Laky-64/http"
)

func main() {
	res, err := http.ExecuteRequest(
		"https://httpbin.org/post",
		http.Method("POST"),
		http.Body([]byte("Hello, World!")),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
