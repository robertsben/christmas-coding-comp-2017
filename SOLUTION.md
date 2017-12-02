# Solution
## Background
We know that the desks with the most presents will be the ones visited by the most elves.
This is easily translated to:
> The number with the most factors receives the most presents.

Of course, as any half-baked mathematician with a knack for googling would know, this means
finding the [number with the most prime factors](http://mathcentral.uregina.ca/QQ/database/QQ.09.03/kristi1.html).

The next challenge is to work out how the desk number will relate to the number of presents.
Given that any composite number `x` can be represented as:
```
x = p1^c1 * p2^c2 * ... * py^cy
```
and that the sum of a composite numbers divisors can be represented as:
```
S(x) = (p1^(c1+1)/p1-1) * (p2^(c2+1)/p2-1)* ... * (py^(cy+1)/py-1)
```
we have a base formulaic relationship to work with.

## Solution

One option is simply to iterate desks (`d`), computing the prime factors of `d`, and then calculating `S(d)`
until we find one `S(d)>PRESENTS`.

That, however, has quite a large complexity (`O(n * p * q)` where `n= desks`, `p= complexity to calculate primes`, 
`q= complexity to calculate sums of primes`) of quite a large `n`.

One way to speed this up is to create a cache of primes, their exponents, and their related sums.
Once we have an array of primes `p1,p2,...pn`, their exponents `p1^0,p1^1,p1^2,...,p1^y, ... ,py^0,py^1,py^2,...,py^y`,
and their sums `S(p1^0),S(p1^1),S(p1^2),...,S(p1^y), ... ,S(py^0),S(py^1),S(py^2),...,S(py^y)`.

Once this is computed and stored, you can find the highest exponent of each prime that divides into each `n`
and use the corollary sum to calculate the presents without needing to repeat the calculation of presents.

The only trouble is, you're still relying on repeated division of every desk `d` from `1->x` where `S(x) > PRESENTS`.
So how can we minimise the number of computed desks?

### Optimisations

Firstly, we can make an assumption that `x` will be even. I don't have a strict mathematical proof of this, but
you can be sure that any odd number will be preceeded by an even one that is divisible by some power of 2, the 
smallest prime. The likelihood is that this means the even number will be more divisible (ergo, have more prime factors)
than the next odd one. From a bit of light observation, this seems to hold, which means we half the number of 
desks we need to calculate.

Secondly, we can attempt to pick a more useful starting point than 1. This is quite tricky, but purely through
observation with no mathematical proof (but a lot of hope) we can make the assumption that `S(x)/5 > x`. This
drastically reduces the number of desks `n` that we have to iterate through.
