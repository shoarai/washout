package sample_washout_test

import (
	"testing"

	"github.com/shoarai/washout"

	"github.com/shoarai/washout/sample_washout"
)

func TestSampleWashoutFilter_inputZero(t *testing.T) {
	expecteds := []washout.Position{
		{Z: -0.95660051821368697},
		{Z: -1.8659615046637352},
		{Z: -2.729832571637687},
		{Z: -3.5499057310185567}}

	const interval = 10
	wash := sample_washout.New(interval)

	for _, expected := range expecteds {
		actual := wash.Filter(0, 0, 0, 0, 0, 0)
		if !isEqualPosition(actual, expected) {
			t.Errorf("wash.Filter() = \n%v,\n wants \n%v", actual, expected)
		}
	}
}

func isEqualPosition(p1, p2 washout.Position) bool {
	return !(p1.X != p2.X ||
		p1.Y != p2.Y ||
		p1.Z != p2.Z ||
		p1.AngleX != p2.AngleX ||
		p1.AngleY != p2.AngleY ||
		p1.AngleZ != p2.AngleZ)
}
