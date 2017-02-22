package vector

//Add adds two vectors elementwise
func Add(v1, v2 []float64) []float64 {
	// if len(v1) != len(v2) {
	// 	panic(fmt.Sprintf("len(v1) != len(v2) [%d, %d]", len(v1), len(v2)))
	// }
	result := make([]float64, len(v1))

	for i, x := range v1 {
		result[i] = x + v2[i]
	}
	return result
}
