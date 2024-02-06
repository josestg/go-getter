# go-getter

go-getter is a tool to generate getter methods for struct fields.

## Install

```shell
go install github.com/josestg/go-getter@latest
```

## Example


```shell
go-getter -typ User -src examples/user.go -out examples/user.gen.go -ptr=false
```


or using `go:generate` directive


```go
// Path: examples/todo.go

package examples

//go:generate go-getter -typ Todo -src $GOFILE -out ./todo.gen.go
type Todo struct {
    title     string
    completed bool
}
```

```shell
go generate ./...
```
