(ns sieve.core
  (:gen-class))

; check if n is prime, given all primes <= n
(defn isprime [n ps]
  (not-any? #(zero? (mod n %)) ps))

; find all primes up to limit
(defn primes
  ([limit] (primes limit 2 (sorted-set)))
  ([limit n ps] (cond
    (> n limit) ps
    (isprime n ps) (recur limit (inc n) (conj ps n))
    true (recur limit (inc n) ps))))

(defn -main
  "Count the number of primes"
  [& args]
  (println (count (primes 100000))))
