// Copyright © 2017 shoarai

// Package washout provides washout filters.
// The washout filter processes a vehicle motion to
// produce a simulator motion to simulate the motion.
package washout

import (
	"math"

	"github.com/shoarai/washout/internal/integral"
	. "github.com/shoarai/washout/internal/vector"
)

// An Washout is a washout filter.
type Washout struct {
	translationHPFs *[3]Filter
	rotationLPFs    *[2]Filter
	rotationHPFs    *[3]Filter

	translationDoubleIntegrals *[3]integral.Integral
	rotationIntegrals          *[3]integral.Integral

	translationScale, rotationScale float64
	gravityVector                   Vector
}

// WashoutFilter is a washout filter.
type WashoutFilter interface {
	Filter(accelerationX, accelerationY, accelerationZ,
		angularVelocityX, angularVelocityY, angularVelocityZ float64) Position
}

type Position struct {
	X, Y, Z, AngleX, AngleY, AngleZ float64
}

type Filter interface {
	Filter(float64) float64
}

// New creates a new washout
func New(
	translationHPFs *[3]Filter,
	rotationLPFs *[2]Filter,
	rotationHPFs *[3]Filter,
	intervalMs uint) *Washout {

	var double uint = 2
	translationDoubleIntegrals := [3]integral.Integral{
		integral.NewMulti(intervalMs, double),
		integral.NewMulti(intervalMs, double),
		integral.NewMulti(intervalMs, double),
	}
	rotationIntegrals := [3]integral.Integral{
		integral.New(intervalMs),
		integral.New(intervalMs),
		integral.New(intervalMs),
	}

	return &Washout{
		translationHPFs:            translationHPFs,
		rotationLPFs:               rotationLPFs,
		rotationHPFs:               rotationHPFs,
		translationDoubleIntegrals: &translationDoubleIntegrals,
		rotationIntegrals:          &rotationIntegrals,
		translationScale:           1,
		rotationScale:              1}
}

// gravity_mm is the acceleration of gravity.
const gravity_mm = 9.80665 * 1000

// Filter processes a vehicle motion to produce a simulator position which simulate the motion
// Filter receive accelarations in meters per second,
// and angular velocities in radians per second as arguments.
// Filter returns X Y Z in meters, and AngleX, AngleY, AngleZ in radians
func (w *Washout) Filter(
	accelerationX, accelerationY, accelerationZ,
	angularVelocityX, angularVelocityY, angularVelocityZ float64) Position {

	scaledAcceralation := Vector{
		accelerationX, accelerationY, accelerationZ,
	}.Multi(w.translationScale)
	scaledAngVel := Vector{
		angularVelocityX, angularVelocityY, angularVelocityZ,
	}.Multi(w.rotationScale)

	displacement := w.translationFilter(scaledAcceralation)

	tiltAngle := w.tiltFilter(scaledAcceralation)
	rotationAngle := w.rotationFilter(scaledAngVel)
	angle := tiltAngle.Plus(rotationAngle)

	w.gravityVector = w.calculateGravity(angle)

	return Position{
		displacement.X, displacement.Y, displacement.Z,
		angle.X, angle.Y, angle.Z}
}

func (w *Washout) translationFilter(acceralation Vector) Vector {
	acce := acceralation.Plus(w.gravityVector)
	acce.Z -= gravity_mm
	acce = Vector{
		w.translationHPFs[0].Filter(acce.X),
		w.translationHPFs[1].Filter(acce.Y),
		w.translationHPFs[2].Filter(acce.Z),
	}

	return Vector{
		w.translationDoubleIntegrals[0].Integrate(acce.X),
		w.translationDoubleIntegrals[1].Integrate(acce.Y),
		w.translationDoubleIntegrals[2].Integrate(acce.Z),
	}
}

func (w *Washout) tiltFilter(acceralation Vector) Vector {
	filteredAx := w.rotationLPFs[0].Filter(acceralation.X)
	filteredAy := w.rotationLPFs[1].Filter(acceralation.Y)

	// TODO: Check if asin returns NaN
	// math.IsNaN(math.Asin(x)

	// Convert low pass filtered accerarations to tilt angles
	return Vector{
		-math.Asin(filteredAy / gravity_mm),
		math.Asin(filteredAx / gravity_mm),
		0}
}

func (w *Washout) rotationFilter(scaledAngVel Vector) Vector {
	filteredAngVelX := w.rotationHPFs[0].Filter(scaledAngVel.X)
	filteredAngVelY := w.rotationHPFs[1].Filter(scaledAngVel.Y)
	filteredAngVelZ := w.rotationHPFs[2].Filter(scaledAngVel.Z)

	return Vector{
		w.rotationIntegrals[0].Integrate(filteredAngVelX),
		w.rotationIntegrals[1].Integrate(filteredAngVelY),
		w.rotationIntegrals[2].Integrate(filteredAngVelZ),
	}
}

// calculateGravity Calculates gravity vector in vehicle coordinate
func (w *Washout) calculateGravity(angle Vector) Vector {
	sphi := math.Sin(angle.X)
	cphi := math.Cos(angle.X)
	ssit := math.Sin(angle.Y)
	csit := math.Cos(angle.Y)

	return Vector{
		-ssit,
		sphi * csit,
		cphi * csit,
	}.Multi(gravity_mm)
}
