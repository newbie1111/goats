package numeric

import (
	"goats/utils"
)

func Factorial[NNI utils.NonNegativeInteger, Z utils.Zahl](n NNI) Z {
	if n == 0 {
		return 1
	} else {
		res, _ := Product(ArithmeticProgression(Z(1), n, Z(1))...)
		return res
	}
}

func IsPrimeByTestDivision[NNI utils.NonNegativeInteger](n NNI) bool {
	switch {
	case n < 2:
		return false
	case n == 2:
		return true
	case n != 2 && n%2 == 0:
		return false
	default:
		for i := NNI(3); i*i <= n; i += 2 {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
}
