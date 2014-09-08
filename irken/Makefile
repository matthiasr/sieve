all: primes

%: %.scm
	irken $< || true # compiler always returns 14

clean:
	rm -f primes.c primes
