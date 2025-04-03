package main

import (
	"fmt"

	"github.com/masterofn0n3/masterofn0n3/nes-emulator/pkg/logicgate"
)

func main() {
	a := logicgate.Input{Value: 1}
	// b := logicgate.Input{Value: 0}
	sel := logicgate.Input{Value: 1}

	out1, out2 := logicgate.DMUX(&a, &sel)

	fmt.Println(out1.Eval(), out2.Eval())
}
