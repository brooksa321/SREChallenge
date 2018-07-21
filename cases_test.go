package main

var fahrenheitTests = []struct {
	input    float64
	target   string
	expected float64
}{
	{84.2, "rankine", 543.9},
	{84.2, "celsius", 29},
	{84.2, "kelvin", 302.1},
}

var celsiusTests = []struct {
	input    float64
	target   string
	expected float64
}{
	{84.2, "rankine", 643.2},
	{84.2, "fahrenheit", 183.6},
	{84.2, "kelvin", 357.3},
}

var kelvinTests = []struct {
	input    float64
	target   string
	expected float64
}{
	{84.2, "rankine", 151.6},
	{84.2, "celsius", -188.9},
	{84.2, "fahrenheit", -308.1},
}

var rankineTests = []struct {
	input    float64
	target   string
	expected float64
}{
	{84.2, "fahrenheit", -375.5},
	{84.2, "celsius", -226.4},
	{84.2, "kelvin", 46.8},
}
