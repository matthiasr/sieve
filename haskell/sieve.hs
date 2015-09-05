primes :: [Integer]

primes =
  nextPrime [2] 3
  where
    nextPrime ps n = 
      if isPrime ps n
      then n : nextPrime (n : ps) ( succ n )
      else nextPrime ps (succ n)
      where
        isPrime ps n =
          all (\ x -> (mod n x) /= 0 ) $ takeWhile (< n) ps

main = print $ last $ take 10000 primes
