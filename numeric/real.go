package numeric

import (
	"errors"
	"goats/utils"
)

func Product[R utils.Real](values ...R) (R, error) {
	if len(values) == 0 {
		return R(0), errors.New(badLength)
	} else {
		product := R(1)

		for _, v := range values {
			product *= v
		}

		return product, nil
	}
}
