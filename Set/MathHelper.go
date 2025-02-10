package Set

// Combination calculates the number of ways to choose r items from n items.
// C(n,r)=(nr)=n!(r!(n−r)!)
func Combination(n, r int) int {
	var b = Factorial(n) / (Factorial(r) * Factorial(n-r))
	return b
}

// Factorial calculates the factorial of a number.
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	var r = n * Factorial(n-1)
	return r
}
