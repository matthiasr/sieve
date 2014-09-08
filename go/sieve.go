/* Erathostenes' Sieve
 * using channels all the way to the bottom
 */

package main

import "fmt"
import "time"

/* buffer size for all channels
 * should trade higher memory usage for less scheduling overhead
 */
const bufsize int = 1000
const max uint64 = 1000000

/* feed all numbers from 0 to MaxUint64 to out */
func numbers(out chan uint64, max uint64) {
	var i uint64

	for i = 0; i <= max; i++ {
		out <- i
	}
	close(out)
}

/* read from in, filter out all multiples of p
 * 
 * first number not filtered is guaranteed to be a prime (if the numbers coming in are pre-filtered with all primes <= p)
 * so feed it to output and set up a new filter goroutine on it
 * feed all further candidates to this new filter
 */
func filter_primes(p uint64, in chan uint64, out chan uint64) {
	var n uint64
	var ok bool
	var next chan uint64

	for {
		n, ok = <-in
		if (n % p) != 0 {
			if next == nil { /* found next prime - set up next filter */
				next = make(chan uint64, bufsize)
				go filter_primes(n, next, out)
				out <- n
			} else { /* feed candidates to next filter */
				next <- n
			}
		}
		/* no more numbers coming - close channels and leave */
		if !ok {
			if next != nil {
				close(next)
			} else {
				close(out)
			}
			return
		}
	}
}

/* helper to set up what filter_primes expects */
func Sieve(out chan uint64, max uint64) {
	in := make(chan uint64, bufsize)

	/* get feed of numbers ... */
	go numbers(in, max)
	/* but drop 0 and 1 */
	for i := uint64(0); i < 2; i = <-in {
	}

	/* 2 is a prime, use it as a starting point */
	out <- 2
	go filter_primes(2, in, out)
}

/* C-Style sieve */
func CSieve(out chan uint64, max uint64) {
	primes := make([]uint64, max)
	for i := 0; i < len(primes); i++ {
		primes[i] = uint64(i + 2)
	}

	for i := 0; i < len(primes); i++ {
		if primes[i] == 0 {
			continue
		}
		for j := i + 1; j < len(primes); j++ {
			//            fmt.Println(i, j, primes[i], primes[j])
			if primes[j] == 0 {
				continue
			}
			if primes[i] != 0 && primes[j]%primes[i] == 0 {
				primes[j] = 0
			}
		}
		out <- primes[i]
	}
	close(out)
}

func main() {
	out := make(chan uint64, bufsize)

	fmt.Println("max =", max, ", bufsize = ", bufsize)

	now := time.Now()
	go Sieve(out, max)
	/* print primes until interrupted or reached MaxUint64 */
	for n := range out {
		n = n
	}
	fmt.Println("Go-Style: ", time.Since(now))

	out = make(chan uint64, bufsize)

	now = time.Now()
	go CSieve(out, max)
	for n := range out {
		n = n
	}
	fmt.Println("C-Style: ", time.Since(now))
}
