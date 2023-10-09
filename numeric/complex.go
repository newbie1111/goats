package numeric

import (
	"errors"
	"goats/utils"
)

func Sum[C utils.Complex](values ...C) (C, error) {
	if len(values) == 0 {
		return 0, errors.New(badLength)
	} else {
		var sum C

		for _, v := range values {
			sum += v
		}

		return sum, nil
	}
}
