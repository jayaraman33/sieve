package sieve

func isPrime(n int) bool {
	for m := 2; m < n; m++ {
		if n%m == 0 {
			return false
		}
	}
	return true
}

func Sieve(num int) []int {
	primes := []int{}
	for n := 2; n <= num; n++ {
		if isPrime(n) {
			primes = append(primes, n)
		}
	}
	return primes
}
