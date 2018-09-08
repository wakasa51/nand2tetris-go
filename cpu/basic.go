package cpu

func Nand(a, b int) int {
	if a == 1 && b == 1 {
		return 0
	} else {
		return 1
	}
}

func Not(a int) int {
	return Nand(a, a)
}

func And(a, b int) int {
	return Not(Nand(a, b))
}

func Or(a, b int) int {
	return Nand(Not(a), Not(b))
}

func Xor(a, b int) int {
	firstNand := Nand(a, b)
	return Nand(Nand(a, firstNand), Nand(firstNand, b))
}

func Mux(a, b, sel int) int {
	a_input := Not(sel)
	b_input := sel
	return Nand(Nand(a, a_input), Nand(b_input, b))
}

func DMux(in, sel int) (a, b int) {
	firstNand := Nand(in, sel)
	a = Not(Nand(in, firstNand))
	b = Not(firstNand)
	return a, b
}

func Not16(inputs [16]int) (outputs [16]int) {
	for i, in := range inputs {
		outputs[i] = Not(in)
	}
	return outputs
}

func And16(a, b [16]int) (outputs [16]int) {
	for i := 0; i < 16; i++ {
		outputs[i] = And(a[i], b[i])
	}
	return outputs
}

func Or16(a, b [16]int) (outputs [16]int) {
	for i := 0; i < 16; i++ {
		outputs[i] = Or(a[i], b[i])
	}
	return outputs
}

func Mux16(a, b [16]int, sel int) (outputs [16]int) {
	for i := 0; i < 16; i++ {
		outputs[i] = Mux(a[i], b[i], sel)
	}
	return outputs
}

func Or8Way(inputs [8]int) (output int) {
	output = 0
	for _, input := range inputs {
		output = Or(input, output)
	}
	return output
}

func Mux4Way16(a, b, c, d [16]int, sel [2]int) (outputs [16]int) {
	outputs = Mux16(Mux16(a, b, sel[0]), Mux16(c, d, sel[0]), sel[1])
	return outputs
}

func Mux8Way16(a, b, c, d, e, f, g, h [16]int, sel [3]int) (outputs [16]int) {
	inter_sel := [2]int{sel[0], sel[1]}
	outputs = Mux16(Mux4Way16(a, b, c, d, inter_sel), Mux4Way16(e, f, g, h, inter_sel), sel[2])
	return outputs
}

func DMux4Way(in int, sel [2]int) (a, b, c, d int) {
	inter_in1, inter_in2 := DMux(in, sel[1])
	a, b = DMux(inter_in1, sel[0])
	c, d = DMux(inter_in2, sel[0])
	return a, b, c, d
}

func DMux8Way(in int, sel [3]int) (a, b, c, d, e, f, g, h int) {
	inter_in1, inter_in2 := DMux(in, sel[2])
	inter_sel := [2]int{sel[0], sel[1]}
	a, b, c, d = DMux4Way(inter_in1, inter_sel)
	e, f, g, h = DMux4Way(inter_in2, inter_sel)
	return a, b, c, d, e, f, g, h
}
