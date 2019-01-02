// Copyright © 2018 shoarai

package washout_test

import (
	"testing"

	"github.com/shoarai/washout"
)

func TestPlus(t *testing.T) {
	v := washout.Vector{1, 2, 3}
	input := washout.Vector{10, 20, 30}
	actual := v.Plus(input)
	expected := washout.Vector{11, 22, 33}
	if actual != expected {
		t.Errorf("\nvector.Plus(%v) = \n%v,\n wants \n%v", input, actual, expected)
	}
}

func TestMulti(t *testing.T) {
	v := washout.Vector{1, 2, 3}
	input := 10.0
	actual := v.Multi(input)
	expected := washout.Vector{10, 20, 30}
	if actual != expected {
		t.Errorf("\nvector.Multi(%v) = \n%v,\n wants \n%v", input, actual, expected)
	}
}
