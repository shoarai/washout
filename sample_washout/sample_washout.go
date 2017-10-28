package sample_washout

import (
	"github.com/shoarai/washout"
)

func New(interval uint) *washout.Washout {
	const cutoffFrequencyForHighPass = 2.5                           // ωn
	const cutoffFrequencyForLowPass = 2 * cutoffFrequencyForHighPass // ωLP
	const dampingRatio = 1                                           // ζLP

	translationHPFs := [3]washout.Filter{}
	for i := range translationHPFs {
		translationHPFs[i] = washout.Filter(&TranslationHighPassFilter{
			SamplingTime:    interval,
			CutoffFrequency: cutoffFrequencyForHighPass})
	}
	rotationLPFs := [2]washout.Filter{}
	for i := range rotationLPFs {
		rotationLPFs[i] = washout.Filter(&RotationLowPassFilter{
			SamplingTime:    interval,
			CutoffFrequency: cutoffFrequencyForLowPass,
			DampingRatio:    dampingRatio})
	}
	rotationHPFs := [3]washout.Filter{}
	for i := range rotationHPFs {
		rotationHPFs[i] = washout.Filter(&RotationHighPassFilter{
			SamplingTime:    interval,
			CutoffFrequency: cutoffFrequencyForHighPass})
	}

	return washout.New(
		&translationHPFs, &rotationLPFs, &rotationHPFs, interval)
}
