package sample_washout

import (
	"github.com/shoarai/washout"
)

func New(interval uint) *washout.Washout {
	const breakFrequencyForHighPass = 2.5                          // ωn
	const breakFrequencyForLowPass = 2 * breakFrequencyForHighPass // ωLP
	const dampingRatio = 1                                         // ζLP

	translationHPFs := [3]washout.Filter{}
	for i := range translationHPFs {
		translationHPFs[i] = washout.Filter(&TranslationHighPassFilter{
			SamplingTime:    interval,
			CutoffFrequency: breakFrequencyForHighPass})
	}
	rotationLPFs := [2]washout.Filter{}
	for i := range rotationLPFs {
		rotationLPFs[i] = washout.Filter(&RotationLowPassFilter{
			SamplingTime:    interval,
			CutoffFrequency: breakFrequencyForLowPass,
			DampingRatio:    dampingRatio})
	}
	rotationHPFs := [3]washout.Filter{}
	for i := range rotationHPFs {
		rotationHPFs[i] = washout.Filter(&RotationHighPassFilter{
			SamplingTime:    interval,
			CutoffFrequency: breakFrequencyForHighPass})
	}

	return washout.New(
		&translationHPFs, &rotationLPFs, &rotationHPFs, interval)
}
