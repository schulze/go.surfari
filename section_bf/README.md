Searching for a section over a \F_p of the elliptic surface

	y^2 = x^3 + t^7x + 1

by brute-force search.

Profiling shows that compiler optimizations don't help here. Most
of the runtime is spend in functions from the flint2 library.

If called as 

	search_bf N

the coefficient of deg_x + 1 is set to N; use this to search in
parallel.
