// Enumeration of singular fibre configurations for K3 surfaces
// with trivial lattice of rank 19, a section and given discriminant.
//
// For the computation of the discriminants see Shioda's "On the
// Mordell-Weil lattices".
package main

import (
	nt "github.com/schulze/go.surfari/frac"
	// "big"
	"fmt"
	"strconv"
	"flag"
)

var D, R int
func init() {
	flag.IntVar(&D, "D", 35, "discriminant of the lattice we want to construct")
	flag.IntVar(&R, "R", 17, "rank of the lattice we want to construct")
}

// Fibre represents the root lattice of a singular fibre.
type Fibre struct {
	e int // Euler number
	r int // rank
	d int // discriminant
	name string // [ADE]_n
	contr []*nt.Frac // contribution of components to the height of sections
}

// Config represents a configuration of singular fibres of ADE-type.
type Config []*Fibre

// The root lattices of singular fibres.
var (
	E6 = NewE(6)
	E7 = NewE(7)
	E8 = NewE(8)
	An = make([]*Fibre, 18)
	Dn = make([]*Fibre, 14)
)

func init() {
	for i := 0; i < 18; i++ {
		An[i] = NewA(i)
	}
	for i := 0; i < 14; i++ {
		Dn[i] = NewD(i)
	}
}

// NewA returns a fibre of type A_n.
func NewA(n int) *Fibre {
	i := strconv.Itoa(n)
	contr := make([]*nt.Frac, (n+1)/2)  // A_n fibres are symmetric
	for i, _ := range contr {
		contr[i] = nt.NewFrac(i*((n+1)-i), n+1)
	}
	return &Fibre{n+1, n, n+1, "A_" + i, contr}
}

// NewD returns a fibre of type D_n, with n >= 4.
func NewD(n int) *Fibre {
	//TODO(fs) should check d >= 4
	i := strconv.Itoa(n)
	contr := make([]*nt.Frac, 4)
	contr[0] = nt.NewFrac(0, 1)
	contr[1] = nt.NewFrac(1, 1)
	contr[2] = nt.NewFrac(n, 4)
	contr[3] = nt.NewFrac(n, 4)

	return &Fibre{n+2, n, 4, "D_" + i, contr}
}

// NewE returns a fibre of type E_n, with n=6,7 or 8.
func NewE(n int) *Fibre {
	i := strconv.Itoa(n)
	// TODO(fs) These are wrong.
	var contr []*nt.Frac
	switch n {
	case 6: contr = []*nt.Frac{nt.NewFrac(0,1), nt.NewFrac(4, 3)}
	case 7: contr = []*nt.Frac{nt.NewFrac(0,1), nt.NewFrac(3, 2)}
	case 8: contr = []*nt.Frac{nt.NewFrac(0,1)}
	}
	return &Fibre{2+n, n, 9-n, "E_" + i, contr}
}

func (f Fibre) String() string {
	return f.name
}

// Disc returns the discriminant of a given root lattice.
func (f Fibre) Disc() (d int) {
	return f.d
}

// Contrib returns a slice of correction terms for the height of sections.
func (f Fibre) Contrib() []*nt.Frac {
	return f.contr
}

// Disc returns the discriminant of a configuration of singular fibres.
func (c Config) Disc() (d int) {
	d = 1
	for _, v := range c {
		d *= v.d
	}
	return d
}

// Euler returns the Euler number of a surface with this configuration of singular fibres.
func (c Config) Euler() (e int) {
	for _, v := range c {
		e += v.e
	}
	return e
}

// Rank returns the rank of the lattice spanned by the configuration of singular fibres.
func (c Config) Rank() (r int) {
	for _, v := range c {
		r += v.r
	}
	return r
}

// String returns the names of the root lattices in a configuration of singular fibres.
func (c Config) String() (s string) {
	for i, v := range c {
		if i > 0 {
			s += ", "
		}
		s += v.String()
	}
	return s
}

// WalkHeights prints all possible ways a section may meet the given configuration
// to produce a lattice with discriminant d.
func (c Config) WalkHeights (d int) {
	c.walkHeightsIter(nt.NewFrac(0,1), c, []int{}, nt.NewFrac(d,1))
}

func (c Config) walkHeightsIter (contr *nt.Frac, rest Config, inter []int, goal *nt.Frac) {
	if len(rest) == 0 {
		four := nt.NewFrac(4,1)
		t := nt.NewFrac(0,1)
		t.Sub(four, contr)
		t.Mul(t, nt.NewFrac(c.Disc(),1))
		if t.Equal(goal) {
			fmt.Println(c, "with inter=", inter, "contr=", contr, "d_T=", c.Disc(), "d_NS=", t)
		}
	} else {
		for i, v := range rest[0].Contrib() {
			c.walkHeightsIter(nt.NewFrac(0,1).Add(contr, v), rest[1:], append(inter, i), goal)
		}
	}
}

// WalkConfigs calls WalkHeights for all fibre configurations with rank <= r.
func WalkConfigs(d, r int, c Config, fibers []*Fibre) {
	if c.Euler() > 24 { // TODO(fs) It is enough to check for r?
		return
	}
	if c.Rank() <= r {
		c.WalkHeights(d)
	}
	if c.Rank() < r {
		for i, v := range fibers {
			WalkConfigs(d, r, append(c, v), fibers[i:])
		}
	}
}

func main() {
	flag.Parse()

	// We search for a configuration with one section that gives discriminant D.
	// We only look for configurations with rank = R and e < 24.
	fibres := []*Fibre{}
	fibres = append(fibres, NewE(6))
	fibres = append(fibres, NewE(7))
	fibres = append(fibres, NewE(8))
	for i := 1; i < 18; i++ {
		fibres = append(fibres, NewA(i))
	}
	for i := 4; i < 14; i++ {
		fibres = append(fibres, NewD(i))
	}
	//for _, v := range fibres {
	//	fmt.Println(v, v.Contrib())
	//}
	for i, v := range fibres {
		WalkConfigs(D, R, Config{v}, fibres[i:])
	}
}
