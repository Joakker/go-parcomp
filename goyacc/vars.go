package main

import "math"

type ValueTable map[string]float64

var Variables = ValueTable{}

func (v ValueTable) Get(name string) float64 {
	if num, ok := v[name]; ok {
		return num
	} else {
		return math.NaN()
	}
}
