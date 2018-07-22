package main

//Confirm fahrenheit conversions are working
var fahrenheitTests = []struct {
	input    float64
	target   string
	expected float64
}{
	{84.2, "rankine", 543.9},
	{84.2, "celsius", 29},
	{84.2, "kelvin", 302.1},
}

//Confirm celsius conversions are working
var celsiusTests = []struct {
	input    float64
	target   string
	expected float64
}{
	{84.2, "rankine", 643.2},
	{84.2, "fahrenheit", 183.6},
	{84.2, "kelvin", 357.3},
}

//Confirm kelvin conversions are working
var kelvinTests = []struct {
	input    float64
	target   string
	expected float64
}{
	{84.2, "rankine", 151.6},
	{84.2, "celsius", -188.9},
	{84.2, "fahrenheit", -308.1},
}

//Confirms rankine conversions are working
var rankineTests = []struct {
	input    float64
	target   string
	expected float64
}{
	{84.2, "fahrenheit", -375.5},
	{84.2, "celsius", -226.4},
	{84.2, "kelvin", 46.8},
}

//Confirms string to float conversion and rounding is working
var strToFloatTest = []struct {
	input    string
	expected float64
}{
	{"45.667", 45.7},
	{"45.644", 45.6},
	{"45.67", 45.7},
	{"45.64", 45.6},
	{"45", 45},
}
