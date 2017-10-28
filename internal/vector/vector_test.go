// Copyright © 2017 shoarai

package vector_test

import (
	"testing"

	. "github.com/shoarai/washout/internal/vector"
)

func TestPlus(t *testing.T) {
	v := Vector{1, 2, 3}
	input := Vector{10, 20, 30}
	actual := v.Plus(input)
	expected := Vector{11, 22, 33}
	if actual != expected {
		t.Errorf("\nvector.Plus(%v) = \n%v,\n wants \n%v", input, actual, expected)
	}
}

func TestMulti(t *testing.T) {
	v := Vector{1, 2, 3}
	input := 10.0
	actual := v.Multi(input)
	expected := Vector{10, 20, 30}
	if actual != expected {
		t.Errorf("\nvector.Multi(%v) = \n%v,\n wants \n%v", input, actual, expected)
	}
}
