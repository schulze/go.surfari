all: section_bf

section_bf: section_bf.c
	gcc -Wall -Wextra -pedantic  -ansi -std=c99 -static -o section_bf -I/usr/local/include section_bf.c -L/usr/local/lib -lflint -lmpfr -lmpir -lc 

clean:
	rm -rf section_bf section_bf.o

.PHONY: all clean 
