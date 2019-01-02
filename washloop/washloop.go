// Copyright © 2018 shoarai

package washloop

import (
	"sync"
	"time"

	"github.com/shoarai/washout"
)

// A Motion is accelerations and angular velocities in 3D axis.
type Motion struct {
	Acceleration    washout.Vector
	AngularVelocity washout.Vector
}

// A WashoutInterface is a interface of washout.
type WashoutInterface interface {
	Filter(
		accelerationX, accelerationY, accelerationZ,
		angularVelocityX, angularVelocityY, angularVelocityZ float64) washout.Position
}

// A WashoutLoop is a loop for process of washout.
type WashoutLoop struct {
	interval uint

	washout  WashoutInterface
	motion   Motion
	position washout.Position

	stopCh chan struct{}

	motionMutex   *sync.Mutex
	positionMutex *sync.Mutex
}

// NewWashLoop creates new washout loop.
func NewWashLoop(washout WashoutInterface, interval uint) *WashoutLoop {
	w := WashoutLoop{}
	w.stopCh = make(chan struct{})
	w.interval = interval
	w.init(washout)
	return &w
}

// Start starts a loop of process.
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

// Stop stops a loop of process.
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

func (w *WashoutLoop) init(washout WashoutInterface) {
	w.washout = washout
	w.motionMutex = new(sync.Mutex)
	w.positionMutex = new(sync.Mutex)
}

// SetMotion sets a motion used as input of washout.
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

// GetPosition gets a position as output of washout.
func (w *WashoutLoop) GetPosition() washout.Position {
	w.positionMutex.Lock()
	defer w.positionMutex.Unlock()

	return w.position
}
