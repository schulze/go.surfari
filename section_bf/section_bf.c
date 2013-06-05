#include <stdio.h>
#include <flint/flint.h>
#include <flint/nmod_poly.h>

const int prime = 43;
const int deg_x = 3;
const int len_x = 4;

int leadingCoeff = -1;
nmod_poly_t rhs, xt7, y;

// Based on Algorithm H in Knuth's 'Generating all n-tuples'.
void generate(int(*fn)(const nmod_poly_t x)) {
	int focus[len_x+1];
	int orient[len_x];
	for (int i = 0; i < len_x; i++) {
		focus[i] = i;
		orient[i] = 1;
	}
	focus[len_x] = len_x;

	nmod_poly_t x;
	nmod_poly_init(x, prime);

	if (leadingCoeff != -1) {
		printf("%d\n", leadingCoeff);
		nmod_poly_set_coeff_ui(x, deg_x + 1, leadingCoeff);
	}		
	for (;;) {
		(*fn)(x);
		
		int j = focus[0];
		focus[0] = 0;
		if (j == len_x) {
			return;
		}
		nmod_poly_set_coeff_ui(x, j, nmod_poly_get_coeff_ui(x, j) + orient[j]);
		if (nmod_poly_get_coeff_ui(x, j) == 0 || nmod_poly_get_coeff_ui(x, j) == (mp_limb_t)(prime-1)) {
			orient[j] = -orient[j];
			focus[j] = focus[j+1];
			focus[j+1] = j + 1;
		}
	}
	nmod_poly_clear(x);
	return;
}

int check_section(const nmod_poly_t x) {
	nmod_poly_pow(rhs, x, 3); // rhs = x^3
	nmod_poly_shift_left(xt7, x, 7);
	nmod_poly_add(rhs, rhs, xt7); // rhs = x^3 + t^7*x
	nmod_poly_set_coeff_ui(rhs, 0, nmod_poly_get_coeff_ui(rhs, 0) + 1); // rhs = x^3 + t^7*x + 1

	if (nmod_poly_sqrt(y, rhs) == 1) {
		printf("x = "); nmod_poly_print(x); printf("\n");
		printf("y = "); nmod_poly_print(y); printf("\n");
	}	
	return 0;
}

int main(int argc, char *argv[]) {
	if (argc > 2) {
		printf("Usage: section_bf [leadingCoeff]\n");
	}
	if (argc == 2) {
		leadingCoeff = atoi(argv[1]);
	}
	
	nmod_poly_init(rhs, prime);
	nmod_poly_init(xt7, prime);
	nmod_poly_init(y, prime);

	generate(check_section);

	nmod_poly_clear(rhs);
	nmod_poly_clear(xt7);	
	nmod_poly_clear(y);
	return 0;
}