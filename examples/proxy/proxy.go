package main

import (
	"fmt"
	"github.com/Laky-64/http"
)

func main() {
	res, err := http.ExecuteRequest(
		"https://ipinfo.io",
		http.Proxy("socks5://127.0.0.1:9050"),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
