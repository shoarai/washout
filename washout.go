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
	gravityVector              Vector
}

// A Filter is a filter returns an output from an input.
type Filter interface {
	Filter(float64) float64
}

// An Position is a position of simulator.
type Position struct {
	X, Y, Z, AngleX, AngleY, AngleZ float64
}

// gravityMM is the acceleration of gravity.
const gravityMM = 9.80665 * 1000

// NewWashout creates a new washout filter.
func NewWashout(
	translationHighPassFilters *[3]Filter,
	translationLowPassFilters *[2]Filter,
	rotationHighPassFilters *[3]Filter, interval uint) *Washout {
	const double = 2
	translationDoubleIntegrals := [3]integral.Integral{
		*integral.NewMulti(interval, double),
		*integral.NewMulti(interval, double),
		*integral.NewMulti(interval, double)}
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
// The filter receives vehicle's accelarations in meters per square second,
// and vehicle's angular velocities in radians per second.
// Then the filter returns simulator's displacements in X, Y, Z-axis in meters
// and simulator's angles in X, Y, Z-axis in radians.
func (w *Washout) Filter(
	accelerationX, accelerationY, accelerationZ,
	angularVelocityX, angularVelocityY, angularVelocityZ float64) Position {

	scaledAcceralation := Vector{
		accelerationX, accelerationY, accelerationZ,
	}.Multi(w.TranslationScale)
	scaledAngVel := Vector{
		angularVelocityX, angularVelocityY, angularVelocityZ,
	}.Multi(w.RotationScale)

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
	acce.Z -= gravityMM
	acce = Vector{
		w.translationHighPassFilters[0].Filter(acce.X),
		w.translationHighPassFilters[1].Filter(acce.Y),
		w.translationHighPassFilters[2].Filter(acce.Z)}

	return Vector{
		w.translationDoubleIntegrals[0].Integrate(acce.X),
		w.translationDoubleIntegrals[1].Integrate(acce.Y),
		w.translationDoubleIntegrals[2].Integrate(acce.Z)}
}

func (w *Washout) tiltFilter(acceralation Vector) Vector {
	filteredAx := w.translationLowPassFilters[0].Filter(acceralation.X)
	filteredAy := w.translationLowPassFilters[1].Filter(acceralation.Y)

	// TODO: Check if asin returns NaN
	// math.IsNaN(math.Asin(x)

	// Convert low pass filtered accerarations to tilt angles
	return Vector{
		math.Asin(filteredAy / gravityMM),
		-math.Asin(filteredAx / gravityMM),
		0}
}

// rotationFilter returns the simulator angle to simulate
func (w *Washout) rotationFilter(scaledAngVel Vector) Vector {
	filteredAngVel := Vector{
		w.rotationHighPassFilters[0].Filter(scaledAngVel.X),
		w.rotationHighPassFilters[1].Filter(scaledAngVel.Y),
		w.rotationHighPassFilters[2].Filter(scaledAngVel.Z)}

	return Vector{
		w.rotationIntegrals[0].Integrate(filteredAngVel.X),
		w.rotationIntegrals[1].Integrate(filteredAngVel.Y),
		w.rotationIntegrals[2].Integrate(filteredAngVel.Z)}
}

// calculateGravity calculates gravity in the simulator coordinate.
func (w *Washout) calculateGravity(angle Vector) Vector {
	sphi := math.Sin(angle.X)
	cphi := math.Cos(angle.X)
	ssit := math.Sin(angle.Y)
	csit := math.Cos(angle.Y)

	return Vector{
		-ssit,
		sphi * csit,
		cphi * csit,
	}.Multi(gravityMM)
}
