(defn isprime [n [p & ps]]
  (cond
    (= (mod n p)) false
    (= (Nil ps)) true
    true (isprime n ps)))
