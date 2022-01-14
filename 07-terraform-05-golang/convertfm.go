package main

import (
	"math"
	"strconv"
)

func convert_fm(input float64) string {
	output := input * 0.3048
	input_str := strconv.FormatFloat(input, 'f', 0, 64)
	output_str := strconv.FormatFloat(math.Ceil(output*10000)/10000, 'f', 4, 64)
	res := input_str + " футов равно " + output_str + " метров."
	return res
}
