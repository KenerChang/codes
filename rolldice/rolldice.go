package rolldice

import (
	"math"
	"math/big"
)

func Compute(n int) *big.Int {
	// assume number order does matter
	// use DP to solve this problem
	// f(x) = 2^x-1 if 1 <= x <= 6
	// f(x) = f(x-1) + f(x-2) + f(x-3) + f(x-4) + f(x-5) + f(x-6) if x > 6
	// the result is exponentail to 2
	// so we use bigInt to handle
	result := make([]*big.Int, n)
	for i := 0; i < n; i++ {
		if i < 6 {
			num := int64(math.Pow(2, float64(i)))
			result[i] = big.NewInt(num)
		} else {
			num := big.NewInt(0)
			num.Add(result[i-1], result[i-2])
			num.Add(num, result[i-3])
			num.Add(num, result[i-4])
			num.Add(num, result[i-5])
			num.Add(num, result[i-6])
			result[i] = num
		}
	}
	return result[n-1]
}