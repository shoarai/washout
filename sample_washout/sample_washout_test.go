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

func TestSampleWashoutFilter_translation(t *testing.T) {
	expecteds := []washout.Position{{
		9.7546105776558483e-05,
		9.7546105776558483e-05,
		-0.95660051821368697,
		6.0661285721471104e-08,
		-6.0661285721471104e-08,
		0,
	}, {
		0.00028787925430374527,
		0.00028787925430374527,
		-1.8659615046637388,
		2.973882543906308e-07,
		-2.973882543906308e-07,
		0,
	}, {
		0.00056658435227004719,
		0.00056658435227004719,
		-2.7298325716377785,
		7.5352070621663112e-07,
		-7.5352070621663112e-07,
		0,
	}, {
		0.00092961765022037905,
		0.00092961765022037883,
		-3.5499057310192748,
		1.4070899278669649e-06,
		-1.4070899278669649e-06,
		0,
	}}

	const interval = 10
	wash := sample_washout.NewWashout(interval)

	for _, expected := range expecteds {
		actual := wash.Filter(1, 1, 0, 0, 0, 0)
		if !isEqualPosition(actual, expected) {
			t.Errorf("wash.Filter() = \n%v,\n wants \n%v", actual, expected)
		}
	}
}

func TestSampleWashoutFilter(t *testing.T) {
	expecteds := []washout.Position{{
		9.7546105776558483e-05,
		0.00019509221155311697,
		-0.95630787989635735,
		0.03950629416207762,
		0.049382655388096997,
		0.059259259259259268,
	}, {
		-0.046932454786220816,
		0.038311495122082491,
		-1.8670097081063179,
		0.078037479397755558,
		0.097545808388304081,
		0.11705532693187017,
	}, {
		-0.18470696036282949,
		0.14896101157535874,
		-2.739307908549077,
		0.11561772784485065,
		0.14451952248359157,
		0.17342433120515735,
	}, {
		-0.4533183430661698,
		0.36352141196027377,
		-3.5833910838269443,
		0.15227048632148069,
		0.19033318308710331,
		0.22840150821243743,
	}}

	const interval = 10
	wash := sample_washout.New(interval)

	for _, expected := range expecteds {
		actual := wash.Filter(1, 2, 3, 4, 5, 6)
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
