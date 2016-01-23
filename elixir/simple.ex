defmodule SimpleSieve do
  defp isprime(n, primes) when is_list(primes) and is_integer(n) do
    !Enum.any?(primes, &(rem(n,&1) == 0))
  end

  defp primes(maxn, n, ps) when is_integer(n) and is_integer(maxn) and n <= maxn and n > 0 do
    cond do
      n == maxn -> ps
      isprime n, ps -> primes maxn, n+1, [n|ps]
      true -> primes maxn, n+1, ps
    end
  end

  def primes(maxn) when is_integer(maxn) and maxn >= 2 do
    primes(maxn,2,[])
  end
end
