// Copyright © 2017 shoarai

package jaxfilter

const cutoffFrequencyForHighPass = 2.5                           // ωn
const cutoffFrequencyForLowPass = 2 * cutoffFrequencyForHighPass // ωLP
const dampingRatio = 1                                           // ζLP

func NewTranslationHighPassFilter(interval uint) *TranslationHighPassFilter {
	return &TranslationHighPassFilter{
		SamplingTime:    interval,
		CutoffFrequency: cutoffFrequencyForHighPass}
}

func NewTranslationHighPassFilters(interval uint) *[3]TranslationHighPassFilter {
	return &[3]TranslationHighPassFilter{
		*NewTranslationHighPassFilter(interval),
		*NewTranslationHighPassFilter(interval),
		*NewTranslationHighPassFilter(interval),
	}
}

func NewRotationLowPassFilter(interval uint) *RotationLowPassFilter {
	return &RotationLowPassFilter{
		SamplingTime:    interval,
		CutoffFrequency: cutoffFrequencyForLowPass,
		DampingRatio:    dampingRatio}
}

func NewRotationLowPassFilters(interval uint) *[2]RotationLowPassFilter {
	return &[2]RotationLowPassFilter{
		*NewRotationLowPassFilter(interval),
		*NewRotationLowPassFilter(interval),
	}
}

func NewRotationHighPassFilter(interval uint) *RotationHighPassFilter {
	return &RotationHighPassFilter{
		SamplingTime:    interval,
		CutoffFrequency: cutoffFrequencyForHighPass}
}

func NewRotationHighPassFilters(interval uint) *[3]RotationHighPassFilter {
	return &[3]RotationHighPassFilter{
		*NewRotationHighPassFilter(interval),
		*NewRotationHighPassFilter(interval),
		*NewRotationHighPassFilter(interval),
	}
}
