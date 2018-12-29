// Copyright © 2018 shoarai

package washloop

import (
	"sync"
	"time"

	"github.com/shoarai/washout"
)

type Vector struct {
	X, Y, Z float64
}

type Motion struct {
	Acceleration    Vector
	AngularVelocity Vector
}

type WashoutLoop struct {
	interval uint

	washout  washout.WashoutInterface
	motion   Motion
	position washout.Position

	stopCh chan struct{}

	motionMutex   *sync.Mutex
	positionMutex *sync.Mutex
}

func NewWashoutLoop(washout washout.WashoutInterface, interval uint) *WashoutLoop {
	w := WashoutLoop{}
	w.stopCh = make(chan struct{})
	w.interval = interval
	w.init(washout)
	return &w
}

func (w *WashoutLoop) Start() {
	interval := time.Duration(w.interval) * time.Millisecond
	ticker := time.NewTicker(interval)
	w.filter()
	for {
		select {
		case <-ticker.C:
			w.filter()
		case <-w.stopCh:
			ticker.Stop()
			return
		}
	}
}

func (w *WashoutLoop) Stop() {
	close(w.stopCh)
}

func (w *WashoutLoop) filter() {
	motion := w.getMotion()
	position := w.washout.Filter(
		motion.Acceleration.X,
		motion.Acceleration.Y,
		motion.Acceleration.Z,
		motion.AngularVelocity.X,
		motion.AngularVelocity.Y,
		motion.AngularVelocity.Z,
	)
	w.setPosition(position)
}

func (w *WashoutLoop) init(washout washout.WashoutInterface) {
	w.washout = washout
	w.motionMutex = new(sync.Mutex)
	w.positionMutex = new(sync.Mutex)
}

func (w *WashoutLoop) SetMotion(motion Motion) {
	w.motionMutex.Lock()
	defer w.motionMutex.Unlock()

	w.motion = motion
}

func (w *WashoutLoop) getMotion() Motion {
	w.motionMutex.Lock()
	defer w.motionMutex.Unlock()

	return w.motion
}

func (w *WashoutLoop) setPosition(position washout.Position) {
	w.positionMutex.Lock()
	defer w.positionMutex.Unlock()

	w.position = position
}

func (w *WashoutLoop) GetPosition() washout.Position {
	w.positionMutex.Lock()
	defer w.positionMutex.Unlock()

	return w.position
}
