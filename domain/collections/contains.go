package collections

// Contains search for an element in a slice
func Contains[T comparable](el T, registers []T) bool {
	for _, register := range registers {
		if el == register {
			return true
		}
	}

	return false
}
