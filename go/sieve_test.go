package main

import "testing"
import "math/rand"
import "math/big"

func TestNumbersMax(t *testing.T) {
    t.Log("Test numbers() honoring maximum correctly")
    out := make(chan uint64)
    var i uint64

    go numbers(out, max)

    i=0
    for n := range out {
        i++
        n = n // silence not-used error
    }
    if i!=max+1 {
        t.Errorf("Got %v numbers, expected %v (bufsize %v)", i, max, bufsize)
    }
}

func TestNumbers(t *testing.T) {
    t.Log("Test numbers() producing correct numbers")
    out := make(chan uint64)
    var n uint64

    go numbers(out, max)

    for i:=uint64(0); i<max; i++ {
        n = <-out
        if i != n {
            t.Errorf("Got %v, expected %v (bufsize %v)", n, i, bufsize)
        }
    }
}

type Siever func(chan uint64, uint64)

func testXSievePrimality(f Siever, t *testing.T) {
    const repeatTests = 4 // how often to repeat the test for each sampled prime
    const sampleSize = 1000 // how many primes to sample
    out := make(chan uint64)

    go f(out, max)

    for n := range out {
        if rand.Int63n(int64(n)) > sampleSize { continue } // skip most tests, otherwise this takes too long
        isprime := big.NewInt(int64(n)).ProbablyPrime(repeatTests)
        if !isprime {
            t.Errorf("Not a prime: %v (bufsize %v)", n, bufsize)
        }
    }
}

// CAUTION: probabilistic test. Will never be falsely negative, may be falsely passing
func TestSievePrimality(t *testing.T) {
    testXSievePrimality(Sieve, t)
}

// CAUTION: probabilistic test. Will never be falsely negative, may be falsely passing
func TestCSievePrimality(t *testing.T) {
    testXSievePrimality(CSieve, t)
}

func TestSieveCrosstest(t *testing.T) {
    t.Log("Test Sieve() and CSieve() equivalency")
    out_sieve := make(chan uint64)
    out_csieve := make(chan uint64)

    var x, y uint64
    var okx, oky bool

    go Sieve(out_sieve, max)
    go CSieve(out_csieve, max)

    for {
        x, okx = <-out_sieve
        y, oky = <-out_csieve
        switch {
            case okx && !oky:
                t.Errorf("CSieve ended prematurely (Sieve: %v, CSieve %v)", x, y)
                return
            case !okx && oky:
                t.Errorf("Sieve ended prematurely (Sieve: %v, CSieve %v)", x, y)
                return
            case !okx && !oky:
                return
        }
        if x!=y { t.Errorf("Got different results:\tSieve: %v\tCSieve:%v", x, y) }
    }
}
