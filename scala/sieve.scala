package sieve

object Sieve {
  def primes(limit: Int) = {
    def _isprime(primes: List[Int], n: Int): Boolean = primes match {
      case Nil => true
      case head :: tail => if ( n%head == 0 ) false else _isprime(tail, n)
    }
    def _primes(primes: List[Int], n: Int): List[Int] =
      if (n > limit)
        primes
      else if (_isprime(primes,n))
        _primes(n :: primes, n+1)
      else
        _primes(primes, n+1)
    _primes(Nil,2)
  }

  def countPrimes(limit: Int) = primes(limit).length

  def main(args: Array[String]) = println(countPrimes(100000))
}
