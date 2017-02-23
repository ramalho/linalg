package vector

import (
	"fmt"
	"math"
)

type operator func(a, b float64) float64

//pairwise applies operator to pairs made from corresponding items in vectors
func pairwise(op operator, v1, v2 []float64) []float64 {
	if len(v1) != len(v2) {
		panic(fmt.Sprintf("Vector lengths are different: %d != %d", len(v1), len(v2)))
	}
	result := make([]float64, len(v1))
	for i, x := range v1 {
		result[i] = op(x, v2[i])
	}
	return result
}

//Add adds two vectors elementwise
func Add(v1, v2 []float64) []float64 {
	return pairwise(func(a, b float64) float64 { return a + b }, v1, v2)
}

//Sub subtracts two vectors elementwise
func Sub(v1, v2 []float64) []float64 {
	return pairwise(func(a, b float64) float64 { return a - b }, v1, v2)
}

//Sum adds vectors
func Sum(vs ...[]float64) []float64 {
	result := make([]float64, len(vs[0]))
	for _, v := range vs {
		result = Add(result, v)
	}
	return result
}

//ScalarMul multiplies a vector by a scalar
func ScalarMul(x float64, v []float64) []float64 {
	result := make([]float64, len(v))
	for i, y := range v {
		result[i] = x * y
	}
	return result
}

//Mean computes the mean of vectors (items are mean of corresponding items)
func Mean(vs ...[]float64) []float64 {
	return ScalarMul(1.0/float64(len(vs)), Sum(vs...))
}

//mul multiplies two vectors elementwise (used by Dot)
func mul(v1, v2 []float64) []float64 {
	return pairwise(func(a, b float64) float64 { return a * b }, v1, v2)
}

//Dot computes the dot product of two vectors
func Dot(v1, v2 []float64) float64 {
	result := 0.0
	for _, x := range mul(v1, v2) {
		result += x
	}
	return result
}

//SumOfSquares computes the dot product v · v
func SumOfSquares(v []float64) float64 {
	return Dot(v, v)
}

//Magnitude of the vector
func Magnitude(v []float64) float64 {
	return math.Sqrt(SumOfSquares(v))
}

//SquaredDistance computes the squared distance of two vectors (used by Distance)
func SquaredDistance(v1, v2 []float64) float64 {
	return SumOfSquares(Sub(v1, v2))
}

//Distance computes the squared distance of two vectors (used by Distance)
func Distance(v1, v2 []float64) float64 {
	return math.Sqrt(SquaredDistance(v1, v2))
}
