// Enumeration of singular fiber configurations for K3 surfaces
// with trivial lattice of rank 19, a section and given discriminant.
//
// For the computation of the discriminants see Shioda's "On the
// Mordell-Weil lattices".
package main

// TODO
// o Should consider also the torsion sections that are possible for each configuration.
//   Then d_NS = disc(trivial lattice) * height(P) / #MW
// o WalkHeights should not print results but return them or pass them down some channel.
// o At some point add discriminant forms computations to compute the transcendental lattice.
//   (c.f. Shimada, Zhang; Classification of extremal elliptic K3 surfaces and fundamental groups of open K3 surfaces)

import (
	nt "github.com/schulze/surfari/frac"
	// "big"
	"flag"
	"fmt"
	"strconv"
)

var (
	D                int // the discriminant of the lattices we are looking for
	R                int // the rank of the lattices we are looking for
	zeroIntersection int // intersection number of section with the zero-section
)

func init() {
	flag.IntVar(&D, "D", 35, "discriminant of the lattice we want to construct")
	flag.IntVar(&R, "R", 17, "rank of the lattice we want to construct")
	flag.IntVar(&zeroIntersection, "intersection", 0, "intersection number of zero section and free section")
}

// Fiber represents the root lattice of a singular fiber.
type Fiber struct {
	e     int        // Euler number
	r     int        // rank
	d     int        // discriminant
	name  string     // [ADE]_n
	contr []*nt.Frac // contribution of components to the height of sections
}

// Config represents a configuration of singular fibers of ADE-type.
type Config []*Fiber

// The root lattices of singular fibers.
var (
	E6 = NewE(6)
	E7 = NewE(7)
	E8 = NewE(8)
	An = make([]*Fiber, 18)
	Dn = make([]*Fiber, 14)
)

func init() {
	for i := 0; i < 18; i++ {
		An[i] = NewA(i)
	}
	for i := 0; i < 14; i++ {
		Dn[i] = NewD(i)
	}
}

// NewA returns a fiber of type A_n.
func NewA(n int) *Fiber {
	i := strconv.Itoa(n)
	contr := make([]*nt.Frac, (n+1)/2+1) // A_n fibers are symmetric
	for i := range contr {
		contr[i] = nt.NewFrac(i*((n+1)-i), n+1)
	}
	return &Fiber{n + 1, n, n + 1, "A_" + i, contr}
}

// NewD returns a fiber of type D_n, with n >= 4.
func NewD(n int) *Fiber {
	//TODO(fs) should check d >= 4
	i := strconv.Itoa(n)
	contr := make([]*nt.Frac, 3) // no contr[3] as the two far components are symmetric
	contr[0] = nt.NewFrac(0, 1)
	contr[1] = nt.NewFrac(1, 1)
	contr[2] = nt.NewFrac(n, 4)

	return &Fiber{n + 2, n, 4, "D_" + i, contr}
}

// NewE returns a fiber of type E_n, with n=6,7 or 8.
func NewE(n int) *Fiber {
	i := strconv.Itoa(n)
	var contr []*nt.Frac
	switch n {
	case 6:
		contr = []*nt.Frac{nt.NewFrac(0, 1), nt.NewFrac(4, 3)}
	case 7:
		contr = []*nt.Frac{nt.NewFrac(0, 1), nt.NewFrac(3, 2)}
	case 8:
		contr = []*nt.Frac{nt.NewFrac(0, 1)}
	}
	return &Fiber{2 + n, n, 9 - n, "E_" + i, contr}
}

func (f Fiber) String() string {
	return f.name
}

// Disc returns the discriminant of a given root lattice.
func (f Fiber) Disc() (d int) {
	return f.d
}

// Contrib returns a slice of correction terms for the height of sections.
func (f Fiber) Contrib() []*nt.Frac {
	return f.contr
}

// Disc returns the discriminant of a configuration of singular fibers.
func (c Config) Disc() (d int) {
	d = 1
	for _, v := range c {
		d *= v.d
	}
	return d
}

// Euler returns the Euler number of a surface with this configuration of singular fibers.
func (c Config) Euler() (e int) {
	for _, v := range c {
		e += v.e
	}
	return e
}

// Rank returns the rank of the lattice spanned by the configuration of singular fibers.
func (c Config) Rank() (r int) {
	for _, v := range c {
		r += v.r
	}
	return r
}

// String returns the names of the root lattices in a configuration of singular fibers.
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
func (c Config) WalkHeights(d int) {
	c.walkHeightsIter(nt.NewFrac(0, 1), c, []int{}, nt.NewFrac(d, 1))
}

func (c Config) walkHeightsIter(contr *nt.Frac, rest Config, inter []int, goal *nt.Frac) {
	if len(rest) == 0 {
		height := nt.NewFrac(4, 1)
		height.Add(height, nt.NewFrac(2*zeroIntersection, 1)) // possible intersection with the zero-section
		height.Sub(height, contr)

		disc := nt.NewFrac(0, 1)
		disc.Mul(height, nt.NewFrac(c.Disc(), 1))
		if disc.Equal(goal) {
			fmt.Println(c, "with inter=", inter, "contr=", contr, "d_T=", c.Disc(), "d_NS=", disc)
		}
	} else {
		for i, v := range rest[0].Contrib() {
			c.walkHeightsIter(nt.NewFrac(0, 1).Add(contr, v), rest[1:], append(inter, i), goal)
		}
	}
}

// WalkConfigs calls WalkHeights for all fiber configurations with rank == r.
func WalkConfigs(d, r int, c Config, fibers []*Fiber) {
	if c.Euler() > 24 {
		return
	}
	if c.Rank() == r {
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
	fibers := []*Fiber{}
	fibers = append(fibers, NewE(6))
	fibers = append(fibers, NewE(7))
	fibers = append(fibers, NewE(8))
	for i := 1; i < 18; i++ {
		fibers = append(fibers, NewA(i))
	}
	for i := 4; i < 14; i++ {
		fibers = append(fibers, NewD(i))
	}
	for i, v := range fibers {
		WalkConfigs(D, R, Config{v}, fibers[i:])
	}
}
