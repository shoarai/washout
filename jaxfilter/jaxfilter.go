// Copyright © 2017 shoarai

package jaxfilter

func square(x float64) float64 {
	return x * x
}

// A TranslationHighPassFilter is a high pass filter for translation.
type TranslationHighPassFilter struct {
	SamplingTime    uint
	CutoffFrequency float64
	inputs          [3]float64
	outputs         [3]float64
}

// Filter passes the high frequency component of a signal.
func (f *TranslationHighPassFilter) Filter(input float64) float64 {
	// T[s]×ωn
	tw := f.CutoffFrequency * float64(f.SamplingTime) / 1000

	f.inputs[0] = input

	// Solve the difference equation
	f.outputs[0] =
		square(2/(tw+2))*(f.inputs[0]-2*f.inputs[1]+f.inputs[2]) -
			2*(tw-2)/(tw+2)*f.outputs[1] -
			square((tw-2)/(tw+2))*f.outputs[2]

	// Delay
	for i := 0; i < 2; i++ {
		f.inputs[2-i] = f.inputs[1-i]
		f.outputs[2-i] = f.outputs[1-i]
	}
	return f.outputs[0]
}

// A TranslationLowPassFilter is a low pass filter for rotation.
type TranslationLowPassFilter struct {
	SamplingTime    uint
	CutoffFrequency float64
	DampingRatio    float64
	inputs          [3]float64
	outputs         [3]float64
}

// Filter passes the low frequency component of a signal.
func (f *TranslationLowPassFilter) Filter(input float64) float64 {
	f.inputs[0] = input

	// (TωLP)^2、4ζωLP*T
	t2w2 := square(f.CutoffFrequency * float64(f.SamplingTime) / 1000)
	dw4T := 4 * f.DampingRatio * f.CutoffFrequency * float64(f.SamplingTime) / 1000

	// Solve the difference equation
	f.outputs[0] =
		(t2w2)/(t2w2+dw4T+4)*
			(f.inputs[0]+2*f.inputs[1]+f.inputs[2]) -
			(1/(t2w2+dw4T+4))*(2*(t2w2-4)*f.outputs[1]+
				(t2w2-dw4T+4)*f.outputs[2])

	// Delay
	for i := 0; i < 2; i++ {
		f.inputs[2-i] = f.inputs[1-i]
		f.outputs[2-i] = f.outputs[1-i]
	}

	return f.outputs[0]

}

// A RotationHighPassFilter is a high pass filter for rotation.
type RotationHighPassFilter struct {
	SamplingTime    uint
	CutoffFrequency float64
	inputs          [2]float64
	outputs         [2]float64
}

// Filter passes the high frequency component of a signal.
func (f *RotationHighPassFilter) Filter(input float64) float64 {
	f.inputs[0] = input

	// Tωn
	tw := f.CutoffFrequency * float64(f.SamplingTime) / 1000

	// Solve the difference equation
	f.outputs[0] = 2/(tw+2)*(f.inputs[0]-f.inputs[1]) -
		(tw-2)/(tw+2)*f.outputs[1]

	// Delay
	f.inputs[1] = f.inputs[0]
	f.outputs[1] = f.outputs[0]

	return f.outputs[0]
}
