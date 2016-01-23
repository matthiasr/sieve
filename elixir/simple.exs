defmodule Sieve do
  def isprime(n, primes) when is_list(primes) and is_integer(n) do
    !Enum.any?(primes, &(rem(n,&1) == 0))
  end

  def primes(maxn, n, ps) when is_integer(n) and is_integer(maxn) and n <= maxn and n > 0 do
    cond do
      n == maxn -> ps
      isprime n, ps -> primes maxn, n+1, [n|ps]
      true -> primes maxn, n+1, ps
    end
  end

  def primes(maxn) when maxn > 0 do
    primes(maxn,3,[2])
  end
end

#Enum.each(Sieve.primes(100000), &IO.puts/1)
IO.puts hd Sieve.primes(100000)
