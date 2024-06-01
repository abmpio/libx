package slicex

// reserve slice
func Reverse[SV any](input []SV) []SV {
	reversedS := make([]SV, len(input))

	for i, v := range input {
		reversedS[len(input)-1-i] = v
	}
	return reversedS
}
