package main

import (
	"math"
	"syscall/js"
)

func add(this js.Value, i []js.Value) interface{} {
	// println(js.ValueOf(i[0].Int() + i[1].Int()).String())
	return js.ValueOf(i[0].Int() + i[1].Int()).String()
}

func subtract(this js.Value, i []js.Value) interface{} {
	// println(js.ValueOf(i[0].Int() - i[1].Int()).String())
	return js.ValueOf(i[0].Int() - i[1].Int()).String()
}

func sqrt(this js.Value, i []js.Value) interface{} {
	a := float64(700000000)
	for i := 0; i < 100; i++ {
		a = math.Sqrt(a)
	}
	return js.ValueOf(a).String()
}

func registerCallbacks() {
	js.Global().Set("wasm_func1", js.FuncOf(add))
	js.Global().Set("wasm_func2", js.FuncOf(subtract))
	js.Global().Set("wasm_func3", js.FuncOf(sqrt))
	js.Global().Set("gameState", js.ValueOf("123"))
}

func main() {
	done := make(chan bool, 0)

	registerCallbacks()

	<-done
}
