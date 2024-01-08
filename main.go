package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{})

	fmt.Println("Hello golang webassembly check")

	js.Global().Set("sayHello", js.FuncOf(sayHello))
	js.Global().Set("add", js.FuncOf(add))

	fmt.Println("done calling from golang")

	<-c
}

func add(val js.Value, values []js.Value) any {
	if values == nil {
		return 0
	}
	return values[0].Int() + values[1].Int()
}

func sayHello(val js.Value, values []js.Value) any {
	fmt.Println("say hello")

	return nil
}
