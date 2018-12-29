// Copyright © 2018 shoarai

package washloop_test

import (
	"testing"
	"time"

	"github.com/shoarai/washout"
	"github.com/shoarai/washout/washloop"
)

var count uint

func TestWashoutLoop(t *testing.T) {
	interval := uint(10)
	loopNumber := uint(5)

	filter := TestWashout{}
	loop := washloop.NewWashoutLoop(filter, interval)

	go func() {
		loop.Start()
	}()

	duration := time.Duration(interval * loopNumber)
	time.Sleep(duration * time.Millisecond)
	loop.Stop()

	if count < loopNumber {
		t.Errorf("Filter() is not processed %v times, want over %v", count, loopNumber)
	}
}

type TestWashout struct{}

func (w TestWashout) Filter(
	accelerationX, accelerationY, accelerationZ,
	angularVelocityX, angularVelocityY, angularVelocityZ float64) washout.Position {
	count++
	return washout.Position{}
}
