package fibonacci
import (
	"math/big"
)

func Compute(n int) *big.Int {
	// use DP to solve fibonacci
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	var n_2, n_1 = big.NewInt(0), big.NewInt(1) // n-2, n-1

	for i := 1; i < n; i++ {
		n_2.Add(n_2, n_1) 
		n_1, n_2 = n_2, n_1 
	}
	return n_1
}