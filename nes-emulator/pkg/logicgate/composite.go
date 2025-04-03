package logicgate

func MUX(a, b, sel Gate) Gate {
	return OR(AND(NOT(sel), a), AND(sel, b))
}

func DMUX(in, sel Gate) (Gate, Gate) {
	return AND(NOT(sel), in), AND(sel, in)
}

func NOT16(inputs [16]Gate) [16]Gate {
	var out [16]Gate
	for i := range inputs {
		out[i] = NOT(inputs[i])
	}

	return out
}

func AND16(inputA [16]Gate, inputB [16]Gate) [16]Gate {
	var out [16]Gate
	for i := 0; i <= 15; i++ {
		out[i] = AND(inputA[i], inputB[i])
	}

	return out
}

func OR16(inputA [16]Gate, inputB [16]Gate) [16]Gate {
	var out [16]Gate
	for i := 0; i <= 15; i++ {
		out[i] = OR(inputA[i], inputB[i])
	}

	return out
}

func MUX16(inputA [16]Gate, inputB [16]Gate, sel Gate) [16]Gate {
	var out [16]Gate
	for i := 0; i <= 15; i++ {
		out[i] = MUX(inputA[i], inputB[i], sel)
	}

	return out
}

func OR8WAY(inputs [8]Gate) Gate {
	out := inputs[0]
	for i := 1; i <= 8; i++ {
		out = OR(out, inputs[i])
	}

	return out
}

func MUX4WAY(a, b, c, d, sel0, sel1 Gate) Gate {
	ab := MUX(a, b, sel0)
	cd := MUX(c, d, sel0)
	return MUX(ab, cd, sel1)
}

func MUX4WAY16(inputs [4][16]Gate, sel [2]Gate) [16]Gate {
	var out [16]Gate
	for i := 0; i < 16; i++ {
		out[i] = MUX4WAY(
			inputs[0][i],
			inputs[1][i],
			inputs[2][i],
			inputs[3][i],
			sel[0],
			sel[1],
		)
	}
	return out
}

func MUX8WAY(
	a, b, c, d, e, f, g, h Gate,
	sel0, sel1, sel2 Gate,
) Gate {
	mux0 := MUX4WAY(a, b, c, d, sel0, sel1)
	mux1 := MUX4WAY(e, f, g, h, sel0, sel1)
	return MUX(mux0, mux1, sel2)
}

func MUX8WAY16(inputs [8][16]Gate, sel [3]Gate) [16]Gate {
	var out [16]Gate
	for i := 0; i < 16; i++ {
		out[i] = MUX8WAY(
			inputs[0][i], inputs[1][i], inputs[2][i], inputs[3][i],
			inputs[4][i], inputs[5][i], inputs[6][i], inputs[7][i],
			sel[0], sel[1], sel[2], // LSB to MSB
		)
	}
	return out
}

func DMUX4WAY(input Gate, sel [2]Gate) [4]Gate {
	sel0 := sel[0] // LSB
	sel1 := sel[1] // MSB

	lowerIn, upperIn := DMUX(input, sel1)

	out0, out1 := DMUX(lowerIn, sel0)
	out2, out3 := DMUX(upperIn, sel0)

	return [4]Gate{out0, out1, out2, out3}
}

func DMUX8WAY(input Gate, sel [3]Gate) [8]Gate {
	sel0 := sel[0] // LSB
	sel1 := sel[1]
	sel2 := sel[2] // MSB

	lower4, upper4 := DMUX(input, sel2)

	lowGroup := DMUX4WAY(lower4, [2]Gate{sel0, sel1})
	highGroup := DMUX4WAY(upper4, [2]Gate{sel0, sel1})

	return [8]Gate{
		lowGroup[0], lowGroup[1], lowGroup[2], lowGroup[3],
		highGroup[0], highGroup[1], highGroup[2], highGroup[3],
	}
}
