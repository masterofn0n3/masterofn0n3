package logicgate

type Gate interface {
	Eval() uint8
}

type BASE struct {
	A uint8
	B uint8
}

// AND gate
type AND struct {
	BASE
}

func (g *AND) Eval() uint8 {
	if g.A > 1 || g.B > 1 {
		panic("Only accept binary")
	}
	return g.A & g.B
}

// OR gate
type OR struct {
	BASE
}

func (g *OR) Eval() uint8 {
	if g.A > 1 || g.B > 1 {
		panic("Only accept binary")
	}
	return g.A | g.B
}

// XOR gate
type XOR struct {
	BASE
}

func (g *XOR) Eval() uint8 {
	if g.A > 1 || g.B > 1 {
		panic("Only accept binary")
	}
	return g.A ^ g.B
}

// NOT gate
type NOT struct {
	Val uint8
}

func (g *NOT) Eval() uint8 {
	if g.Val > 1 {
		panic("Only accept binary")
	}
	// since we're using uint8 which have 8 bits, `^` flips all 8 bits, created unwanted result,
	// so we need `& 1` to mask the `^` result, only left the LEAST SIGNIFICANT BIT
	return ^g.Val & 1
}

type NAND struct {
	BASE
}

func (g *NAND) Eval() uint8 {
	andGate := &AND{g.BASE}

	return (&NOT{Val: andGate.Eval()}).Eval()
}
