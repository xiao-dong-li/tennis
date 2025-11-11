package game

// CloneMatrix returns a deep copy of the given 2D matrix.
func CloneMatrix(src [][]bool) [][]bool {
	dst := make([][]bool, len(src))
	for i := range src {
		dst[i] = append([]bool(nil), src[i]...)
	}
	return dst
}
