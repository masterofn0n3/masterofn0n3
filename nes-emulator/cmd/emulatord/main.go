package main

import (
	"fmt"

	"github.com/masterofn0n3/masterofn0n3/nes-emulator/pkg/logicgate"
)

func main() {
	nandGate := logicgate.NAND{BASE: logicgate.BASE{A: 0, B: 0}}
	fmt.Println(nandGate.Eval())
}
