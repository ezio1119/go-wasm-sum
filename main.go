package main

import (
	"fmt"
	"strconv"
	"syscall/js"
	"time"
)

func calcSum(a, b int) int {
	return a + b
}

func calcSumJs() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			for {
				time.Sleep(time.Second)
				fmt.Println("ループ中")
			}
		}()

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
	js.Global().Set("calcSum", calcSumJs())

	<-make(chan struct{})
}
