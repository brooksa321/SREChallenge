package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//Converts kelvin inputs
func kelvinConversions(temp float64, target string) float64 {
	var kelvConv float64
	switch {
	case target == "celsius":
		kelvConv = temp - 273.15
	case target == "fahrenheit":
		kelvConv = temp*(9/5) - 459.67
	case target == "rankine":
		kelvConv = temp * (9 / 5)
	}

	return kelvConv
}

//Converts celcius inputs
func celsiusConversions(temp float64, target string) float64 {
	var celsConv float64
	switch {
	case target == "kelvin":
		celsConv = temp + 273.15
	case target == "fahrenheit":
		celsConv = temp*(9/5) + 32
	case target == "rankine":
		celsConv = (temp + 273.15) * (9 / 5)
	}

	return celsConv
}

//Converts fahrenheit inputs
func fahrenheitConversions(temp float64, target string) float64 {
	var fahrConv float64
	switch {
	case target == "kelvin":
		fahrConv = (temp + 459.67) * (5 / 9)
	case target == "celsius":
		fahrConv = (temp - 32) * (5 / 9)
	case target == "rankine":
		fahrConv = temp + 459.67
	}

	return fahrConv
}

//Converts rankine inputs
func rankineConversions(temp float64, target string) float64 {
	var rankConv float64
	switch {
	case target == "kelvin":
		rankConv = temp * (5 / 9)
	case target == "fahrenheit":
		rankConv = temp - 459.67
	case target == "celsius":
		rankConv = (temp - 491.67) * (5 / 9)
	}

	return rankConv
}

//Converts temperatures from strings to floats with error validation
func strToFloat(num string) float64 {
	tempTest, err := strconv.ParseFloat(num, 64)
	if err != nil {
		log.Fatalln("invalid")
	}

	roundTemp := fmt.Sprintf("%.1f", tempTest)
	fmt.Println(roundTemp)
	temperature, _ := strconv.ParseFloat(roundTemp, 64)

	return temperature
}

func checkResp(answer float64, response float64) string {
	var check string
	if answer == response {
		check = "correct"
	} else {
		check = "incorrect"
	}
	return check
}

func main() {
	//Allow teacher to enter Input Temperature.
	tempReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Input Temperature: ")
	inputTemp, _ := tempReader.ReadString('\n')
	inputTemp = strings.TrimSuffix(inputTemp, "\n")

	//Allow teacher to input target conversion units
	unitReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Target Units: ")
	targetUnits, _ := unitReader.ReadString('\n')
	targetUnits = strings.ToLower(targetUnits)
	targetUnits = strings.TrimSuffix(targetUnits, "\n")

	//Allow teacher to input Student Resp
	respReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Student Response: ")
	inputResp, _ := respReader.ReadString('\n')
	inputResp = strings.TrimSuffix(inputResp, "\n")

	//Allow Input temperature parameters to be handled independently
	inputSplit := strings.Split(inputTemp, " ")

	//Set input Unit
	inputUnit := strings.ToLower(inputSplit[1])

	temperature := strToFloat(inputSplit[0])
	response := strToFloat(inputResp)

	//Selects which conversion to run based on input unit
	switch {
	case inputUnit == "kelvin":
		fmt.Println(checkResp(kelvinConversions(temperature, targetUnits), response))
	case inputUnit == "celsius":
		fmt.Println(checkResp(celsiusConversions(temperature, targetUnits), response))
	case inputUnit == "fahrenheit":
		fmt.Println(checkResp(fahrenheitConversions(temperature, targetUnits), response))
	case inputUnit == "rankine":
		fmt.Println(checkResp(rankineConversions(temperature, targetUnits), response))
	}
}
