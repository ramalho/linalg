package vector

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		v1   []float64
		v2   []float64
		want []float64
	}{
		{[]float64{1.0, 1.0}, []float64{0.0, 0.0}, []float64{1.0, 1.0}},
	}
	for _, c := range cases {
		got := Add(c.v1, c.v2)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Add(%v, %v) want: %v; got: %v",
				c.v1, c.v2, c.want, got)
		}
	}
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func TestAdd_whenVectorLensAreDifferent(t *testing.T) {
	v1 := []float64{1.0, 2.0}
	v2 := []float64{1.0}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Add(%v, %v) did not panic.", v1, v2)
		}
	}()

	Add(v1, v2)

}
