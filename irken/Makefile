all: run

run: primes
	./primes || true

%: %.scm
	irken $< || true # compiler always returns 14

clean:
	rm -f primes.c primes

.PHONY: all run clean
