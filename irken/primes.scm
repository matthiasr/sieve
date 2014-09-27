(include "lib/core.scm")
(include "lib/pair.scm")

(define isprime?
  n () -> #t
  n (p . ps) -> (if (zero? (remainder n p))
                #f
                (isprime? n ps)))

(define (primes limit n ps)
  (cond
    ((>= n limit) ps)
    ((isprime? n ps) (primes limit (+ n 1) (cons n ps)))
    (else (primes limit (+ n 1) ps))))

(length (primes 100000 2 (list:nil)))
