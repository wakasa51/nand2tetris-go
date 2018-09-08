package cpu

import (
	"testing"
)

func TestNand(t *testing.T) {
	a := []int{0, 0, 1, 1}
	b := []int{0, 1, 0, 1}
	expects := []int{1, 1, 1, 0}

	for i, expect := range expects {
		actual := Nand(a[i], b[i])
		if expect != actual {
			t.Errorf("index: %d expect: %d actual: %d", i, expect, actual)
		}
	}
}

func TestNot(t *testing.T) {
	a := []int{0, 1}
	expects := []int{1, 0}

	for i, expect := range expects {
		actual := Not(a[i])
		if expect != actual {
			t.Errorf("index: %d expect: %d actual: %d", i, expect, actual)
		}
	}
}

func TestAnd(t *testing.T) {
	a := []int{0, 0, 1, 1}
	b := []int{0, 1, 0, 1}
	expects := []int{0, 0, 0, 1}

	for i, expect := range expects {
		actual := And(a[i], b[i])
		if expect != actual {
			t.Errorf("index: %d expect: %d actual: %d", i, expect, actual)
		}
	}
}

func TestOr(t *testing.T) {
	a := []int{0, 0, 1, 1}
	b := []int{0, 1, 0, 1}
	expects := []int{0, 1, 1, 1}

	for i, expect := range expects {
		actual := Or(a[i], b[i])
		if expect != actual {
			t.Errorf("index: %d expect: %d actual: %d", i, expect, actual)
		}
	}
}

func TestXor(t *testing.T) {
	a := []int{0, 0, 1, 1}
	b := []int{0, 1, 0, 1}
	expects := []int{0, 1, 1, 0}

	for i, expect := range expects {
		actual := Xor(a[i], b[i])
		if expect != actual {
			t.Errorf("index: %d expect: %d actual: %d", i, expect, actual)
		}
	}
}

func TestMux(t *testing.T) {
	a := []int{0, 0, 1, 1, 0, 0, 1, 1}
	b := []int{0, 1, 0, 1, 0, 1, 0, 1}
	sel := []int{0, 0, 0, 0, 1, 1, 1, 1}
	expects := []int{0, 0, 1, 1, 0, 1, 0, 1}

	for i, expect := range expects {
		actual := Mux(a[i], b[i], sel[i])
		if expect != actual {
			t.Errorf("index: %d expect: %d actual: %d", i, expect, actual)
		}
	}
}

func TestDMux(t *testing.T) {
	in := []int{1, 1}
	sel := []int{0, 1}
	e_a := []int{1, 0}
	e_b := []int{0, 1}

	for i, s := range sel {
		a_a, a_b := DMux(in[i], s)
		if e_a[i] != a_a {
			t.Errorf("a: index: %d expect: %d actual: %d", i, e_a[i], a_a)
		}
		if e_b[i] != a_b {
			t.Errorf("b: index: %d expect: %d actual: %d", i, e_b[i], a_b)
		}
	}
}

func TestNot16(t *testing.T) {
	a := [16]int{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 1}
	expects := [16]int{1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0}
	actual := Not16(a)

	if expects != actual {
		t.Errorf("expect: %v actual: %v", expects, actual)
	}
}

func TestAnd16(t *testing.T) {
	a := [16]int{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1}
	b := [16]int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	expects := [16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1}
	actual := And16(a, b)

	if expects != actual {
		t.Errorf("expect: %v actual: %v", expects, actual)
	}
}

func TestOr16(t *testing.T) {
	a := [16]int{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1}
	b := [16]int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	expects := [16]int{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	actual := Or16(a, b)

	if expects != actual {
		t.Errorf("expect: %v actual: %v", expects, actual)
	}
}

func TestMux16(t *testing.T) {
	a := [16]int{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1}
	b := [16]int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	sels := []int{0, 1}
	expects := [][16]int{
		{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
	}

	for i, expect := range expects {
		actual := Mux16(a, b, sels[i])
		if expect != actual {
			t.Errorf("expect: %v actual: %v", expect, actual)
		}
	}
}

func TestOr8Way(t *testing.T) {
	include1 := [8]int{0, 0, 0, 0, 1, 1, 1, 1}
	notInclude1 := [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	expect1 := 1
	expect0 := 0
	actual1 := Or8Way(include1)
	actual0 := Or8Way(notInclude1)

	if expect1 != actual1 {
		t.Errorf("expect: %v actual: %v", expect1, actual1)
	}
	if expect0 != actual0 {
		t.Errorf("expect: %v actual: %v", expect0, actual0)
	}
}

func TestMux4Way16(t *testing.T) {
	a := [16]int{1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	b := [16]int{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}
	c := [16]int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0}
	d := [16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1}
	sels := [][2]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}}
	expects := [][16]int{
		{1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1},
	}

	for i, expect := range expects {
		actual := Mux4Way16(a, b, c, d, sels[i])
		if expect != actual {
			t.Errorf("expect: %v actual: %v", expect, actual)
		}
	}
}

func TestMux8Way16(t *testing.T) {
	a := [16]int{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	b := [16]int{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	c := [16]int{0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	d := [16]int{0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}
	e := [16]int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0}
	f := [16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}
	g := [16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0}
	h := [16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1}
	sels := [][3]int{{0, 0, 0}, {1, 0, 0}, {0, 1, 0}, {1, 1, 0}, {0, 0, 1}, {1, 0, 1}, {0, 1, 1}, {1, 1, 1}}
	expects := [][16]int{
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
	}

	for i, expect := range expects {
		actual := Mux8Way16(a, b, c, d, e, f, g, h, sels[i])
		if expect != actual {
			t.Errorf("expect: %v actual: %v", expect, actual)
		}
	}
}

func TestDMux4Way(t *testing.T) {
	in := []int{1, 1, 1, 1}
	sels := [][2]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}}
	expects := [][4]int{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}

	for i, sel := range sels {
		a, b, c, d := DMux4Way(in[i], sel)
		actual := [4]int{a, b, c, d}
		if expects[i] != actual {
			t.Errorf("index: %d expect: %d actual: %d", i, expects[i], actual)
		}
	}
}

func TestDMux8Way(t *testing.T) {
	in := []int{1, 1, 1, 1, 1, 1, 1, 1}
	sel := [][3]int{{0, 0, 0}, {1, 0, 0}, {0, 1, 0}, {1, 1, 0}, {0, 0, 1}, {1, 0, 1}, {0, 1, 1}, {1, 1, 1}}
	expects := [][8]int{
		{1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 1},
	}

	for i, s := range sel {
		a, b, c, d, e, f, g, h := DMux8Way(in[i], s)
		actual := [8]int{a, b, c, d, e, f, g, h}
		if expects[i] != actual {
			t.Errorf("index: %d expect: %d actual: %d", i, expects[i], actual)
		}
	}
}
