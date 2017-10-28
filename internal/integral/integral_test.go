// Copyright © 2017 shoarai

package integral_test

import (
	"testing"

	"github.com/shoarai/washout/internal/integral"
)

func TestIntegral(t *testing.T) {
	expecteds := []float64{
		2, 4, 6, 8,
	}

	var interval uint = 1000
	input := 2.0
	integ := integral.New(interval)

	for _, expected := range expecteds {
		actual := integ.Integrate(input)
		if actual != expected {
			t.Errorf("integ.Integrate(%f) = \n%v,\n wants \n%v", input, actual, expected)
		}
	}
}

func TestMultiIntegral(t *testing.T) {
	expecteds := []float64{
		2, 6, 12, 20,
	}

	var interval uint = 1000
	input := 2.0
	integ := integral.NewMulti(interval, 2)

	for _, expected := range expecteds {
		actual := integ.Integrate(input)
		if actual != expected {
			t.Errorf("integ.Integrate(%f) = \n%v,\n wants \n%v", input, actual, expected)
		}
	}
}
