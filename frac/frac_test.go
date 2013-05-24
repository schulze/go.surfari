package frac

import (
	"testing"
)

// TODO(fs) Use random numbers and compare with math/big.

var (
	x = NewFrac(2,3)
	y = NewFrac(3,2)
	one1 = NewFrac(1,1)
	one6 = NewFrac(6,6)
)

func TestOne(t *testing.T) {
	if ! one1.Equal(one6) {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	var z Frac
	z.Add(x,y)
	if ! z.Equal(NewFrac(13, 6)) {
		t.Fail()
	}
}

func TestSub(t *testing.T) {
	var z Frac
	z.Sub(x,y)
	if ! z.Equal(NewFrac(-5, 6)) {
		t.Fatalf("x = %v instead of -5/6", z)
	}
}

func TestMul(t *testing.T) {
	var z Frac
	z.Mul(x,y)
	if ! z.Equal(one6) {
		t.Fail()
	}
	if ! z.Equal(one1) {
		t.Fail()
	}
}

func TestDiv(t *testing.T) {
	var z Frac
	z.Div(x,y)
	if ! z.Equal(NewFrac(4, 9)) {
		t.Fatalf("x = %v instead of 4/9", z)
	}
}
