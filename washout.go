// Copyright © 2017 shoarai

// Package washout provides washout filters
// to approximately display the sensation of vehicle motions.
package washout

import (
	"math"

	"github.com/shoarai/washout/internal/integral"
	. "github.com/shoarai/washout/internal/vector"
)

// An Washout is a washout filter.
type Washout struct {
	TranslationScale, RotationScale float64

	translationHighPassFilters *[3]Filter
	translationLowPassFilters  *[2]Filter
	rotationHighPassFilters    *[3]Filter

	translationDoubleIntegrals *[3]integral.Integral
	rotationIntegrals          *[3]integral.Integral
	simulatorGravity           Vector
}

// A Filter is a filter returns an output from an input.
type Filter interface {
	Filter(input float64) (output float64)
}

// An Position is a position of simulator.
type Position struct {
	X, Y, Z, AngleX, AngleY, AngleZ float64
}

// gravity is the acceleration of gravity.
const gravity = 9.80665 * 1000

// NewWashout creates a new washout filter.
// interval is the interval of proccessing in milliseconds.
func NewWashout(
	translationHighPassFilters *[3]Filter,
	translationLowPassFilters *[2]Filter,
	rotationHighPassFilters *[3]Filter, interval uint) *Washout {
	// For accelerations to velocities to displacements
	const integralNumber = 2
	translationDoubleIntegrals := [3]integral.Integral{
		*integral.NewMulti(interval, integralNumber),
		*integral.NewMulti(interval, integralNumber),
		*integral.NewMulti(interval, integralNumber)}

	// For angular velocities to angles
	rotationIntegrals := [3]integral.Integral{
		*integral.New(interval),
		*integral.New(interval),
		*integral.New(interval)}

	return &Washout{
		TranslationScale:           1,
		RotationScale:              1,
		translationHighPassFilters: translationHighPassFilters,
		translationLowPassFilters:  translationLowPassFilters,
		rotationHighPassFilters:    rotationHighPassFilters,
		translationDoubleIntegrals: &translationDoubleIntegrals,
		rotationIntegrals:          &rotationIntegrals}
}

// Filter processes vehicle motions to produce simulator positions to simulate the motion.
// The filter receives vehicle's accelerations in meters per square second,
// and vehicle's angular velocities in radians per second.
// Then the filter returns simulator's displacements in X, Y, Z-axis in meters
// and simulator's angles in X, Y, Z-axis in radians.
func (w *Washout) Filter(
	accelerationX, accelerationY, accelerationZ,
	angularVelocityX, angularVelocityY, angularVelocityZ float64) Position {
	scaledAcceleration := Vector{
		accelerationX, accelerationY, accelerationZ,
	}.Multi(w.TranslationScale)
	scaledAngVel := Vector{
		angularVelocityX, angularVelocityY, angularVelocityZ,
	}.Multi(w.RotationScale)

	displacement := w.toSimulatorDisplacement(&scaledAcceleration)

	tiltAngle := w.toSimulatorTilt(&scaledAcceleration)
	rotationAngle := w.toSimulatorRotation(&scaledAngVel)
	angle := tiltAngle.Plus(rotationAngle)

	w.simulatorGravity = w.calculateGravity(&angle)

	return Position{
		displacement.X, displacement.Y, displacement.Z,
		angle.X, angle.Y, angle.Z}
}

func (w *Washout) toSimulatorDisplacement(acceleration *Vector) Vector {
	acce := acceleration.Plus(w.simulatorGravity)
	acce.Z -= gravity
	acce = w.filterVector(w.translationHighPassFilters, &acce)
	return w.integrateVector(w.translationDoubleIntegrals, &acce)
}

func (w *Washout) toSimulatorTilt(acceleration *Vector) Vector {
	filteredAx := w.translationLowPassFilters[0].Filter(acceleration.X)
	filteredAy := w.translationLowPassFilters[1].Filter(acceleration.Y)

	// TODO: Check if asin returns NaN.
	// math.IsNaN(math.Asin(x)

	// Convert low pass filtered accelerations to tilt angles.
	return Vector{
		math.Asin(filteredAy / gravity),
		-math.Asin(filteredAx / gravity),
		0}
}

// toSimulatorRotation returns the simulator angle to simulate.
func (w *Washout) toSimulatorRotation(angVel *Vector) Vector {
	filteredAngVel := w.filterVector(w.rotationHighPassFilters, angVel)
	return w.integrateVector(w.rotationIntegrals, &filteredAngVel)
}

// calculateGravity calculates gravity in the simulator coordinate.
func (w *Washout) calculateGravity(angle *Vector) Vector {
	sinAngleX := math.Sin(angle.X)
	cosAngleX := math.Cos(angle.X)
	sinAngleY := math.Sin(angle.Y)
	cosAngleY := math.Cos(angle.Y)

	return Vector{
		-sinAngleY,
		sinAngleX * cosAngleY,
		cosAngleX * cosAngleY,
	}.Multi(gravity)
}

func (w *Washout) filterVector(filter *[3]Filter, vector *Vector) Vector {
	return Vector{
		filter[0].Filter(vector.X),
		filter[1].Filter(vector.Y),
		filter[2].Filter(vector.Z)}
}

func (w *Washout) integrateVector(integ *[3]integral.Integral, vector *Vector) Vector {
	return Vector{
		integ[0].Integrate(vector.X),
		integ[1].Integrate(vector.Y),
		integ[2].Integrate(vector.Z)}
}
