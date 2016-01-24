defmodule StreamSieve do
  def ismult(n, ps) when is_integer(n) do
    !Enum.all?(
      Stream.take_while(ps, &(&1 < n)),
      &(rem(n,&1) == 0)
    )
  end

  def primes(n, ps) do
    cond do
      ismult(n, ps) -> primes( n+1, Stream.concat(ps,[n]))
      true -> primes(n+1, ps)
    end
  end

  def primes do
    primes(2,[])
  end

  def primes n do
    Enum.take(primes, n)
  end
end
