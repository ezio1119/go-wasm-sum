package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func calcSum(a, b int) int {
	return a + b
}

func calcSumJs() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Printf("%#v\n", args[0].String())
		if len(args) != 2 {
			return "Invalid arguments passed"
		}
		inputA, err := strconv.Atoi(args[0].String())
		if err != nil {
			return err.Error()
		}

		inputB, err := strconv.Atoi(args[1].String())
		if err != nil {
			return err.Error()
		}

		return calcSum(inputA, inputB)
	})
}

func main() {

	<-make(chan struct{})
}
