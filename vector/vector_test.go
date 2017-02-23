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
		v1   []float64
		v2   []float64
		want []float64
	}{
		{[]float64{1.0, 1.0}, []float64{0.0, 0.0}, []float64{1.0, 1.0}},
		{[]float64{1.0, 2.0, 3.0}, []float64{4.0, 5.0, 6.0}, []float64{5.0, 7.0, 9.0}},
	}
	for _, c := range cases {
		got := Add(c.v1, c.v2)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Add(%v, %v) want: %v; got: %v",
				c.v1, c.v2, c.want, got)
		}
	}
}

func TestAdd_panicWhenVectorLengthsDiffer(t *testing.T) {
	v1 := []float64{1.0}
	v2 := []float64{1.0, 2.0}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Add(%v, %v) did not panic.", v1, v2)
		} else if !strings.HasPrefix(fmt.Sprint(r), "Vector lengths are different") {
			t.Errorf("Add(%v, %v): unexpected panic %q", v1, v2, r)
		}
	}()

	Add(v1, v2)
}

func TestSub(t *testing.T) {
	cases := []struct {
		v1   []float64
		v2   []float64
		want []float64
	}{
		{[]float64{1.0, 1.0}, []float64{0.0, 0.0}, []float64{1.0, 1.0}},
		{[]float64{1.0, 2.0, 3.0}, []float64{4.0, 5.0, 6.0}, []float64{-3.0, -3.0, -3.0}},
	}
	for _, c := range cases {
		got := Sub(c.v1, c.v2)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Add(%v, %v) want: %v; got: %v",
				c.v1, c.v2, c.want, got)
		}
	}
}

func TestSum(t *testing.T) {
	cases := []struct {
		vs   [][]float64
		want []float64
	}{
		{[][]float64{[]float64{1.0, 2.0}, []float64{3.0, 4.0}, []float64{5.0, 6.0}},
			[]float64{9.0, 12.0}},
		{[][]float64{[]float64{1.0, 2.0}},
			[]float64{1.0, 2.0}},
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
		x    float64
		v    []float64
		want []float64
	}{
		{2.0, []float64{3.0, 4.0}, []float64{6.0, 8.0}},
		{0.0, []float64{3.0, 4.0, 5.0}, []float64{0.0, 0.0, 0.0}},
		{1.0, []float64{}, []float64{}},
	}
	for _, c := range cases {
		got := ScalarMul(c.x, c.v)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("ScalarMul(%v, %v) want: %v; got: %v",
				c.x, c.v, c.want, got)
		}
	}
}

func TestMean(t *testing.T) {
	cases := []struct {
		vs   [][]float64
		want []float64
	}{
		{[][]float64{[]float64{1.0, 2.0}, []float64{3.0, 4.0}, []float64{5.0, 6.0}},
			[]float64{3.0, 4.0}},
		{[][]float64{[]float64{1.0, 2.0}},
			[]float64{1.0, 2.0}},
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
		v1   []float64
		v2   []float64
		want float64
	}{
		{[]float64{1.0, 1.0}, []float64{0.0, 0.0}, 0.0},
		{[]float64{1.0, 2.0, 3.0}, []float64{4.0, 5.0, 6.0}, 32.0},
	}
	for _, c := range cases {
		got := Dot(c.v1, c.v2)
		if got != c.want {
			t.Errorf("Dot(%v, %v) want: %v; got: %v",
				c.v1, c.v2, c.want, got)
		}
	}
}

func TestSumOfSquares(t *testing.T) {
	cases := []struct {
		v    []float64
		want float64
	}{
		{[]float64{3.0, 4.0}, 25.0},
		{[]float64{1.0, 2.0, 3.0}, 14.0},
	}
	for _, c := range cases {
		got := SumOfSquares(c.v)
		if got != c.want {
			t.Errorf("SumOfSquares(%v) want: %v; got: %v",
				c.v, c.want, got)
		}
	}
}

func TestMagnitude(t *testing.T) {
	cases := []struct {
		v    []float64
		want float64
	}{
		{[]float64{3.0, 4.0}, 5.0},
		{[]float64{1.0, 2.0, 3.0}, math.Sqrt(14.0)},
	}
	for _, c := range cases {
		got := Magnitude(c.v)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Magnitude(%v) want: %v; got: %v",
				c.v, c.want, got)
		}
	}
}

func TestSquaredDistance(t *testing.T) {
	cases := []struct {
		v1   []float64
		v2   []float64
		want float64
	}{
		{[]float64{1.0, 1.0}, []float64{0.0, 0.0}, 2.0},
		{[]float64{3.0, 0.0}, []float64{0.0, 4.0}, 25.0},
		{[]float64{1.0, 2.0, 3.0}, []float64{4.0, 5.0, 6.0}, 27.0},
	}
	for _, c := range cases {
		got := SquaredDistance(c.v1, c.v2)
		if got != c.want {
			t.Errorf("SquaredDistance(%v, %v) want: %v; got: %v",
				c.v1, c.v2, c.want, got)
		}
	}
}

func TestDistance(t *testing.T) {
	cases := []struct {
		v1   []float64
		v2   []float64
		want float64
	}{
		{[]float64{1.0, 1.0}, []float64{1.0, 2.0}, 1.0},
		{[]float64{3.0, 0.0}, []float64{0.0, 4.0}, 5.0},
		{[]float64{1.0, 2.0, 3.0}, []float64{4.0, 5.0, 6.0}, math.Sqrt(27.0)},
	}
	for _, c := range cases {
		got := Distance(c.v1, c.v2)
		if got != c.want {
			t.Errorf("Distance(%v, %v) want: %v; got: %v",
				c.v1, c.v2, c.want, got)
		}
	}
}
