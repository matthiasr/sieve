(ns sieve.core
  (:gen-class))

; check if n is prime, given all primes <= n
(defn isprime [n ps]
  (cond
    (empty? ps) true
    (= 0 (mod n (first ps))) false
    true (recur n (rest ps))))

; find all primes up to limit
(defn primes
  ([limit n ps] (cond
    (> n limit) ps
    (isprime n ps) (recur limit (inc n) (conj ps n))
    true (recur limit (inc n) ps)))
  ([limit] (primes limit 2 (sorted-set))))

(defn -main
  "I don't do a whole lot ... yet."
  [& args]
  (println (count (primes 100000))))
