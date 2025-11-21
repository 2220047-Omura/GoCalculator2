CC = clang

#CFLAGS = -O3 -g -Wall -Wextra -Werror -pedantic -std=c99 \
#CFLAGS = -O3 -g -Wall -Wextra -Werror -std=c99 \
	-I/opt/homebrew/include
#LDFLAGS = -L/opt/homebrew/lib -lmpfi -lmpfr -lgmp
LDFLAGS = -L/opt/homebrew/lib -lmpfi -lmpfr -lgmp -pthread
CFLAGS_OMP = -O3 -g -Wall -Wextra -Werror -std=c99 \
	-I/opt/homebrew/include -fopenmp
#CFLAGS_OMP = -O3 -g -Wall -Wextra -Werror \
	-I/opt/homebrew/include -fopenmp
LDFLAGS_OMP = -L/opt/homebrew/lib -lmpfi -lmpfr -lgmp -fopenmp

#EXE = LU-cgo-PartlyWG LU-cgo-PartlyWG-type1 LU-cgo-PartlyWG-type1-omp
EXE = LU-cgo-PartlyWG-type1 LU-cgo-PartlyWG-type1-omp

target: $(EXE)

LU-cgo-PartlyWG-type1-omp: LU-cgo-PartlyWG-type1-omp.c
	$(CC) $(CFLAGS_OMP) -c $^ -o LU-cgo-PartlyWG-type1-omp.o
	$(CC) $(LDFLAGS_OMP) LU-cgo-PartlyWG-type1-omp.o -o $@

testnest-omp: testnest-omp.c
	$(CC) $(CFLAGS_OMP) -c $^ -o testnest-omp.o
	$(CC) $(LDFLAGS_OMP) testnest-omp.o -o $@

simple: simple.o
	$(CC) $(LDFLAGS) $^ -o $@

simplep: simplep.o
	$(CC) $(LDFLAGS) $^ -o $@

simple_copy: simple_copy.o
	$(CC) $(LDFLAGS) $^ -o $@

.c.o:
	$(CC) $(CFLAGS) -c $<


clean:
	rm -f *.o $(EXE)