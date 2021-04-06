package problem0050

/* ベタな解答 */
// func myPow(x float64, n int) float64 {
// 	result := 1.
// 	if n > 0 {
// 		for i := 0; i < n; i++ {
// 			result *= x
// 		}
// 	} else if n < 0 {
// 		for i := 0; i < -n; i++ {
// 			result /= x
// 		}
// 	}
// 	return result
// }
/* 当然 Time Limit Exceeded */

func myPow(x float64, n int) float64 {
	k := 1.
	if n == 0 {
		return k
	}
	is_positive := n > 0
	if !is_positive {
		n *= -1
	}
	for n > 1 {
		if n%2 == 0 {
			x = x * x
			n = n / 2
		} else {
			k *= x
			x = x * x
			n = (n - 1) / 2
		}
	}
	result := 1.
	if is_positive {
		result *= k * x
	} else {
		result /= k * x
	}
	return result
}
