package main

import (
	"syscall/js"
)

func sum(args []js.Value) {
	var sum int
	for _, val := range args {
		sum += val.Int()
	}
	println(sum)
}

func registerCallbacks() {
	js.Global().Set("sum", js.NewCallback(sum))
}


func main() {
	c := make(chan struct{}, 0)
	println("Hello, WebAssembly!")
	registerCallbacks()
	<-c
}
