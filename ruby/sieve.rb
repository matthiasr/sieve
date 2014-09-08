n_to_check = (2..100).to_a.reverse!
primes = []

until n_to_check.empty?
    i = n_to_check.pop
    n_to_check.delete_if { |j| j % i == 0 }
    primes << i
end

p primes
