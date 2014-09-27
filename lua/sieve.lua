function any(f,list)
  for _,e in ipairs(list) do
    if f(e) then
      return true
    end
  end
  return false
end

primes = {}
for i = 2,100000 do
  if not any(function (p) return (i%p == 0) end, primes) then
    table.insert(primes,i)
  end
end

print(#primes)
