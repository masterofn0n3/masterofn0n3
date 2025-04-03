package main

import (
	"fmt"

	"github.com/masterofn0n3/masterofn0n3/nes-emulator/pkg/logicgate"
)

func main() {
	a := logicgate.Input{Value: 1}
	b := logicgate.Input{Value: 1}
	nand := logicgate.NAND(&a, &b)
	fmt.Println(nand.Eval())
}
