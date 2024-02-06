package examples

//go:generate go-getter -typ Todo -src $GOFILE -out ./todo.gen.go
type Todo struct {
    title     string
    completed bool
}
