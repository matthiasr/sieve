actor Main
  new create(env: Env) =>
    """Start the program"""
    env.out.print("Starting…")
    let s = PrimeSieve.create(2)
    var i: U64 = 3
    while i < 1000000 do
      s.filter(i)
      i = i+1
    end

actor PrimeSieve
  let p: U64
  var next: ( PrimeSieve | None ) = None

  new create(p': U64) =>
    p = p'

  fun candidate(i: U64): Bool =>
    """Checks whether the argument is not a multiple of the prime"""
    ( i % p ) != 0

  be filter(i: U64) =>
    """Filters all multiples of p. Spawns a new filter for the first non-multiple, then feeds all subsequent non-multiples into it."""
    if candidate(i) then
      try
        (next as PrimeSieve).filter(i)
      else
        next = PrimeSieve.create(i)
      end
    end
