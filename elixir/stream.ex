defmodule StreamSieve do
  def isprime(n, ps) when is_integer(n) do
    !Enum.all?(
      Stream.take_while(ps, &(&1 < n)),
      &(rem(n,&1) == 0)
    )
  end

  def prime_generator [ n: n, ps: ps ] do
    cond do
      isprime(n, ps) -> { n, [ n: (n+1), ps: ps ] }
      true -> prime_generator [ n: n+1, ps: ps ]
    end
  end

  def primes do
    Stream.unfold([ n: 3, ps: [2] ], &prime_generator/1)
  end

  def primes n do
    Enum.take(primes, n)
  end
end
