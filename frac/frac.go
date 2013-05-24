// Simplest possible package for arithmetic with fractions of ints
// and the first Go code I ever wrote.
package frac

import (
	"strconv"
)

type Frac struct {
	p, q int
}

func gcd(a, b int) int {
        for b != 0 {
                a, b = b, a%b
        }
    return abs(a) // TODO(fs): This should be checked earlier. See Cohen.
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func lcm(a, b int) int {
	d := gcd(a, b)
	return a*b / d
}

func NewFrac(p, q int) *Frac {
	d := gcd(p, q)
	if q < 0 {
		q = -q
		p = -p
	}
	return &Frac{p/d, q/d}
}

func (z *Frac) Equal(x *Frac) bool {
	if z.p == x.p && z.q == x.q {
		return true
	}
	return false
}

func (z *Frac) String() string {
	return strconv.Itoa(z.p) + "/" + strconv.Itoa(z.q)
}

func reduce(p, q int) (int, int) {
	d := gcd(p, q)
	if q < 0 {
		q = -q
		p = -p
	}
	return p/d, q/d
}	

func (z *Frac) Add(x, y *Frac) *Frac {
	numer := x.p*y.q + y.p*x.q
	denom := x.q * y.q
	z.p, z.q = reduce(numer, denom)
	return z
}

func (z *Frac) Sub(x, y *Frac) *Frac {
	numer := x.p*y.q - y.p*x.q
	denom := x.q * y.q
	z.p, z.q = reduce(numer, denom)
	return z
}

func (z *Frac) Mul(x, y *Frac) *Frac {
	z.p, z.q = reduce(x.p*y.p, x.q*y.q)
	return z
}

func (z *Frac) Div(x, y *Frac) *Frac {
	z.p, z.q = reduce(x.p*y.q, x.q*y.p)
	return z
}
