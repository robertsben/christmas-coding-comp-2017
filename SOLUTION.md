# Solution

I had so many big dreams and great ideas to do with calculating the number of divisors; using caches of 
primes and their exponents; concurrency, parallelism. But no. No, it all boiled down to about 30 lines of
`C`, with one stupidly simple double for loop. 3 weeks of research into the mathematics of factorisation, 
the divisor function, the geometric series behind its growth with increasing `n`, profiling and
optimising `Go`, all to waste. The moral of the story? Keep It Simple, Stupid.