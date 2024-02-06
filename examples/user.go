package examples

type User struct {
    id   int
    age  int
    name string
    pass string `go-getter:"ignored"`
}
