package main

// Source: exercism/problem-specifications
// Commit: 99de15d raindrops: apply "input" policy
// Problem Specifications Version: 1.1.0

var fahrenheitTests = []struct {
	input    float64
	target   string
	expected float64
}{
	{84.2, "rankine", 543.87},
	{84.2, "celsius", 29},
	{84.2, "kelvin", 302.15},
}

var celsiusTests = []struct {
	input    float64
	target   string
	expected float64
}{
	{84.2, "rankine", 643.23},
	{84.2, "celsius", -226.3722},
	{84.2, "kelvin", 46.77778},
}

// var kelvinTests = []struct {
// 	input    float64
// 	target   string
// 	expected float64
// }{
// 	{84.2, "rankine", 543.87},
// 	{84.2, "celsius", 29},
// 	{84.2, "kelvin", 302.15},
// }

// var rankineTests = []struct {
// 	input    float64
// 	target   string
// 	expected float64
// }{
// 	{84.2, "rankine", 543.87},
// 	{84.2, "celsius", 29},
// 	{84.2, "kelvin", 302.15},
// }
