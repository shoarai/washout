// Copyright © 2017 shoarai

package washout_test

import (
	"testing"

	. "github.com/shoarai/washout"
)

func TestWashoutFilter(t *testing.T) {
	washout := newTestWashoutFilter()
	actual := washout.Filter(0, 0, 0, 0, 0, 0)
	expected := Position{Z: -0.9806650000000001}

	if !isEqualPosition(actual, expected) {
		t.Errorf("Filter() = %v, want %v", actual, expected)
	}
}

func isEqualPosition(p1, p2 Position) bool {
	return !(p1.X != p2.X ||
		p1.Y != p2.Y ||
		p1.Z != p2.Z ||
		p1.AngleX != p2.AngleX ||
		p1.AngleY != p2.AngleY ||
		p1.AngleZ != p2.AngleZ)
}

func newTestWashoutFilter() *Washout {
	translationHPFs := [3]Filter{
		&testFilter{1},
		&testFilter{1},
		&testFilter{1},
	}
	rotationLPFs := [2]Filter{
		&testFilter{1},
		&testFilter{1},
	}
	rotationHPFs := [3]Filter{
		&testFilter{1},
		&testFilter{1},
		&testFilter{1},
	}
	interval_ms := uint(10)

	return NewWashout(&translationHPFs, &rotationLPFs, &rotationHPFs, interval_ms)
}

type testFilter struct {
	Id int
}

func (filter *testFilter) Filter(input float64) float64 {
	return input
	// return float64(filter.Id) + val
}
