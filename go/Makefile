all: sieve run test

run: sieve
	./sieve

sieve: sieve.go
	go build $<

test:
	go test

clean:
	rm -f go sieve

.PHONY: all run test clean
