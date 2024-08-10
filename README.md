# http
A small simplification of the standard library net/http library in Go.

## Usage

Use `go get` to download the dependency.

```bash
go get github.com/Laky-64/http@latest
```

Then, `import` it in your Go files:

```go
import "github.com/Laky-64/http"
```

The library is designed to be simple and easy to use. 
Here's an example of a simple request:


### Simple Request
```go
res, err := http.ExecuteRequest(
	"https://httpbin.org/get",
)
if err != nil {
    panic(err)
}
fmt.Println(res)
```
<img src="https://vhs.charm.sh/vhs-1Qsv8thjvA9KpxkDvyvB4.gif" alt="Example of a simple request">

### POST Request

```go
res, err := http.ExecuteRequest(
	"https://httpbin.org/post",
	http.Method("POST"),
	http.Body([]byte("Hello, World!")),
)
if err != nil {
    panic(err)
}
fmt.Println(res)
```
<img src="https://vhs.charm.sh/vhs-1gJR3CtJNcKPiY3r9g4tS8.gif" alt="Example of a post request">

### MultiPart Request
 ```go
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
```
<img src="https://vhs.charm.sh/vhs-1CmVZwbWkBqglhss7Gw09g.gif" alt="Example of a multipart request">