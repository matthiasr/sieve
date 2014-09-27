primes = []

(2..100000).each do |n|
  unless primes.any? { |p| n%p == 0 }
    primes << n
  end
end

p primes.length
