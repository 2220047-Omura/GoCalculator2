

CFLAGS = -O3 -g -Wall -Wextra -Werror -pedantic -std=c99 \
	-I/opt/homebrew/include
LDFLAGS = -L/opt/homebrew/lib -lmpfi -lmpfr -lgmp

simple: simple.o
	$(CC) $(LDFLAGS) $^ -o $@

simplep: simplep.o
	$(CC) $(LDFLAGS) $^ -o $@

.c.o:
	$(CC) $(CFLAGS) -c $<


clean:
	rm -f *.o
