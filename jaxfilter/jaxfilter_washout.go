// Copyright © 2017 shoarai

package jaxfilter

import "github.com/shoarai/washout"

func NewWashout(interval uint) *washout.Washout {
	translationHPFs := [3]washout.Filter{}
	for i := range translationHPFs {
		translationHPFs[i] = NewTranslationHighPassFilter(interval)
	}
	rotationLPFs := [2]washout.Filter{}
	for i := range rotationLPFs {
		rotationLPFs[i] = NewRotationLowPassFilter(interval)
	}
	rotationHPFs := [3]washout.Filter{}
	for i := range rotationHPFs {
		rotationHPFs[i] = NewRotationHighPassFilter(interval)
	}

	return washout.NewWashout(
		&translationHPFs, &rotationLPFs, &rotationHPFs, interval)
}
