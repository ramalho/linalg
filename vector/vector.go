package vector

import (
	"fmt"
	"math"
)

//Vector represents one point in N-dimensional space
type Vector []float64

type operator func(a, b float64) float64

//pairwise applies operator to pairs made from corresponding items in vectors
func pairwise(op operator, v1, v2 Vector) Vector {
	if len(v1) != len(v2) {
		panic(fmt.Sprintf("Vector lengths are different: %d != %d", len(v1), len(v2)))
	}
	result := make(Vector, len(v1))
	for i, x := range v1 {
		result[i] = op(x, v2[i])
	}
	return result
}

//Add adds two vectors elementwise
func (v Vector) Add(v2 Vector) Vector {
	return pairwise(func(a, b float64) float64 { return a + b }, v, v2)
}

//Sub subtracts two vectors elementwise
func (v Vector) Sub(v2 Vector) Vector {
	return pairwise(func(a, b float64) float64 { return a - b }, v, v2)
}

//Sum adds vectors
func Sum(vs ...Vector) Vector {
	result := make(Vector, len(vs[0]))
	for _, v := range vs {
		result = result.Add(v)
	}
	return result
}

//ScalarMul multiplies a vector by a scalar
func (v Vector) ScalarMul(x float64) Vector {
	result := make(Vector, len(v))
	for i, y := range v {
		result[i] = x * y
	}
	return result
}

//Mean computes the mean of vectors (items are mean of corresponding items)
func Mean(vs ...Vector) Vector {
	return Sum(vs...).ScalarMul(1.0 / float64(len(vs)))
}

//mul multiplies two vectors elementwise (used by Dot)
func (v Vector) mul(v2 Vector) Vector {
	return pairwise(func(a, b float64) float64 { return a * b }, v, v2)
}

//Dot computes the dot product of two vectors
func (v Vector) Dot(v2 Vector) float64 {
	result := 0.0
	for _, x := range v.mul(v2) {
		result += x
	}
	return result
}

//SumOfSquares computes the dot product v·v
func (v Vector) SumOfSquares() float64 {
	return v.Dot(v)
}

//Magnitude of the vector
func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.SumOfSquares())
}

//SquaredDistance computes the squared distance of two vectors (used by Distance)
func (v Vector) SquaredDistance(v2 Vector) float64 {
	return v.Sub(v2).SumOfSquares()
}

//Distance computes the distance of two vectors
func (v Vector) Distance(v2 Vector) float64 {
	return math.Sqrt(v.SquaredDistance(v2))
}
