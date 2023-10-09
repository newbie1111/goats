package numeric

import "goats/utils"

func ArithmeticProgression[Z utils.Zahl, NNI utils.NonNegativeInteger](first Z, terms NNI, diff Z) []Z {
	sequence := make([]Z, 0, terms)

	for i := NNI(0); i < terms; i++ {
		sequence = append(sequence, first+diff*Z(i))
	}

	return sequence
}
