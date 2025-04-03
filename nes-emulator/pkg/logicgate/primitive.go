package logicgate

type Gate interface {
	Eval() uint8
}

// Input represents a binary input signal (0 or 1)
type Input struct {
	Value uint8
}

func (i *Input) Eval() uint8 {
	if i.Value > 1 {
		panic("Only binary values (0 or 1) allowed")
	}
	return i.Value
}

func (i *Input) Set(v uint8) {
	i.Value = v
}

// NOT Gate
type not struct {
	In Gate
}

func (g *not) Eval() uint8 {
	in := g.In.Eval()
	if in > 1 {
		panic("Only binary input allowed")
	}
	// since we're using uint8 which have 8 bits, `^` flips all 8 bits, created unwanted result,
	// so we need `& 1` to mask the `^` result, only left the LEAST SIGNIFICANT BIT
	return ^in & 1
}

func NOT(in Gate) Gate {
	return &not{In: in}
}

// AND Gate
type and struct {
	A, B Gate
}

func (g *and) Eval() uint8 {
	a := g.A.Eval()
	b := g.B.Eval()
	return a & b
}

func AND(a, b Gate) Gate {
	return &and{A: a, B: b}
}

// OR Gate
type or struct {
	A, B Gate
}

func (g *or) Eval() uint8 {
	a := g.A.Eval()
	b := g.B.Eval()
	return a | b
}

func OR(a, b Gate) Gate {
	return &or{A: a, B: b}
}

// XOR Gate
type xor struct {
	A, B Gate
}

func (g *xor) Eval() uint8 {
	a := g.A.Eval()
	b := g.B.Eval()
	return a ^ b
}

func XOR(a, b Gate) Gate {
	return &xor{A: a, B: b}
}

// NAND Gate (NOT of AND)
func NAND(a, b Gate) Gate {
	return NOT(AND(a, b))
}
