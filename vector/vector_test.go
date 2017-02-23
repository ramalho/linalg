package vector

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		v1   Vector
		v2   Vector
		want Vector
	}{
		{Vector{1.0, 1.0}, Vector{0.0, 0.0}, Vector{1.0, 1.0}},
		{Vector{1.0, 2.0, 3.0}, Vector{4.0, 5.0, 6.0}, Vector{5.0, 7.0, 9.0}},
	}
	for _, c := range cases {
		got := c.v1.Add(c.v2)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("%v.Add(%v) want: %v; got: %v",
				c.v1, c.v2, c.want, got)
		}
	}
}

func TestAdd_panicWhenVectorLengthsDiffer(t *testing.T) {
	v1 := Vector{1.0}
	v2 := Vector{1.0, 2.0}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%v.Add(%v) did not panic.", v1, v2)
		} else if !strings.HasPrefix(fmt.Sprint(r), "Vector lengths are different") {
			t.Errorf("%v.Add(%v): unexpected panic %q", v1, v2, r)
		}
	}()

	v1.Add(v2)
}

func TestSub(t *testing.T) {
	cases := []struct {
		v1   Vector
		v2   Vector
		want Vector
	}{
		{Vector{1.0, 1.0}, Vector{0.0, 0.0}, Vector{1.0, 1.0}},
		{Vector{1.0, 2.0, 3.0}, Vector{4.0, 5.0, 6.0}, Vector{-3.0, -3.0, -3.0}},
	}
	for _, c := range cases {
		got := c.v1.Sub(c.v2)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("%v.Sub(%v) want: %v; got: %v",
				c.v1, c.v2, c.want, got)
		}
	}
}

func TestSum(t *testing.T) {
	cases := []struct {
		vs   []Vector
		want Vector
	}{
		{[]Vector{Vector{1.0, 2.0}, Vector{3.0, 4.0}, Vector{5.0, 6.0}},
			Vector{9.0, 12.0}},
		{[]Vector{Vector{1.0, 2.0}},
			Vector{1.0, 2.0}},
	}
	for _, c := range cases {
		got := Sum(c.vs...)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Sum(%v) want: %v; got: %v",
				c.vs, c.want, got)
		}
	}
}

func TestScalarMul(t *testing.T) {
	cases := []struct {
		v    Vector
		x    float64
		want Vector
	}{
		{Vector{3.0, 4.0}, 2.0, Vector{6.0, 8.0}},
		{Vector{3.0, 4.0, 5.0}, 0.0, Vector{0.0, 0.0, 0.0}},
		{Vector{}, 1.0, Vector{}},
	}
	for _, c := range cases {
		got := c.v.ScalarMul(c.x)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("%v.ScalarMul(%v) want: %v; got: %v",
				c.v, c.x, c.want, got)
		}
	}
}

func TestMean(t *testing.T) {
	cases := []struct {
		vs   []Vector
		want Vector
	}{
		{[]Vector{Vector{1.0, 2.0}, Vector{3.0, 4.0}, Vector{5.0, 6.0}},
			Vector{3.0, 4.0}},
		{[]Vector{Vector{1.0, 2.0}},
			Vector{1.0, 2.0}},
	}
	for _, c := range cases {
		got := Mean(c.vs...)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Mean(%v) want: %v; got: %v",
				c.vs, c.want, got)
		}
	}
}

func TestDot(t *testing.T) {
	cases := []struct {
		v1   Vector
		v2   Vector
		want float64
	}{
		{Vector{1.0, 1.0}, Vector{0.0, 0.0}, 0.0},
		{Vector{1.0, 2.0, 3.0}, Vector{4.0, 5.0, 6.0}, 32.0},
	}
	for _, c := range cases {
		got := c.v1.Dot(c.v2)
		if got != c.want {
			t.Errorf("%v.Dot(%v) want: %v; got: %v",
				c.v1, c.v2, c.want, got)
		}
	}
}

func TestSumOfSquares(t *testing.T) {
	cases := []struct {
		v    Vector
		want float64
	}{
		{Vector{3.0, 4.0}, 25.0},
		{Vector{1.0, 2.0, 3.0}, 14.0},
	}
	for _, c := range cases {
		got := c.v.SumOfSquares()
		if got != c.want {
			t.Errorf("%v.SumOfSquares() want: %v; got: %v",
				c.v, c.want, got)
		}
	}
}

func TestMagnitude(t *testing.T) {
	cases := []struct {
		v    Vector
		want float64
	}{
		{Vector{3.0, 4.0}, 5.0},
		{Vector{1.0, 2.0, 3.0}, math.Sqrt(14.0)},
	}
	for _, c := range cases {
		got := c.v.Magnitude()
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("%v.Magnitude() want: %v; got: %v",
				c.v, c.want, got)
		}
	}
}

func TestSquaredDistance(t *testing.T) {
	cases := []struct {
		v1   Vector
		v2   Vector
		want float64
	}{
		{Vector{1.0, 1.0}, Vector{0.0, 0.0}, 2.0},
		{Vector{3.0, 0.0}, Vector{0.0, 4.0}, 25.0},
		{Vector{1.0, 2.0, 3.0}, Vector{4.0, 5.0, 6.0}, 27.0},
	}
	for _, c := range cases {
		got := c.v1.SquaredDistance(c.v2)
		if got != c.want {
			t.Errorf("%v.SquaredDistance(%v) want: %v; got: %v",
				c.v1, c.v2, c.want, got)
		}
	}
}

func TestDistance(t *testing.T) {
	cases := []struct {
		v1   Vector
		v2   Vector
		want float64
	}{
		{Vector{1.0, 1.0}, Vector{1.0, 2.0}, 1.0},
		{Vector{3.0, 0.0}, Vector{0.0, 4.0}, 5.0},
		{Vector{1.0, 2.0, 3.0}, Vector{4.0, 5.0, 6.0}, math.Sqrt(27.0)},
	}
	for _, c := range cases {
		got := c.v1.Distance(c.v2)
		if got != c.want {
			t.Errorf("%v.Distance(%v) want: %v; got: %v",
				c.v1, c.v2, c.want, got)
		}
	}
}
