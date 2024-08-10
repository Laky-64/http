package main

import (
	"fmt"
	"github.com/Laky-64/http"
	"github.com/Laky-64/http/types"
)

func main() {
	res, err := http.ExecuteRequest(
		"https://httpbin.org/post",
		http.Method("POST"),
		http.MultiPartForm(
			map[string]string{
				"key": "value",
			},
			map[string]types.FileDescriptor{
				"file": {
					FileName: "file.txt",
					Content:  []byte("Hello, World!"),
				},
			},
		),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
