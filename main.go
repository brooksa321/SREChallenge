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
		kelvConv = temp*9/5 - 459.67
	case target == "rankine":
		kelvConv = temp * 9 / 5
	}

	kelvConv = roundFloat(kelvConv)
	fmt.Printf("%v kelvin is %v %v \n", temp, kelvConv, target)
	return kelvConv
}

//Converts celcius inputs
func celsiusConversions(temp float64, target string) float64 {
	var celsConv float64
	switch {
	case target == "kelvin":
		celsConv = temp + 273.15
	case target == "fahrenheit":
		celsConv = temp*9/5 + 32
	case target == "rankine":
		celsConv = (temp + 273.15) * 9 / 5
	}

	celsConv = roundFloat(celsConv)
	fmt.Printf("%v celsius is %v %v \n", temp, celsConv, target)
	return celsConv
}

//Converts fahrenheit inputs
func fahrenheitConversions(temp float64, target string) float64 {
	var fahrConv float64
	switch {
	case target == "kelvin":
		fahrConv = (temp + 459.67) * 5 / 9
	case target == "celsius":
		fahrConv = (temp - 32) * 5 / 9
	case target == "rankine":
		fahrConv = temp + 459.67
	}

	fahrConv = roundFloat(fahrConv)
	fmt.Printf("%v fahrenheit is %v %v \n", temp, fahrConv, target)
	return fahrConv
}

//Converts rankine inputs
func rankineConversions(temp float64, target string) float64 {
	var rankConv float64
	switch {
	case target == "kelvin":
		rankConv = temp * 5 / 9
	case target == "fahrenheit":
		rankConv = temp - 459.67
	case target == "celsius":
		rankConv = (temp - 491.67) * 5 / 9
	}

	rankConv = roundFloat(rankConv)
	fmt.Printf("%v rankine is %v %v \n", temp, rankConv, target)
	return rankConv
}

//Converts temperatures from strings to floats with error validation
func strToFloat(num string) float64 {
	tempTest, err := strconv.ParseFloat(num, 64)
	if err != nil {
		log.Fatalln("INVALID")
	}

	temperature := roundFloat(tempTest)

	return temperature
}

//Validates Student Response to Authorotative Answer
func checkResp(answer float64, response float64) string {
	var check string
	if answer == response {
		check = "CORRECT"
	} else {
		check = "INCORRECT"
	}
	return check
}

//Rounds temperatures to ones place
func roundFloat(temp float64) float64 {
	rounded := fmt.Sprintf("%.1f", temp)
	newFloat, _ := strconv.ParseFloat(rounded, 64)

	return newFloat
}

//Confirms a valid unit for conversion was entered
func unitValidation(unit string) {
	if unit != "rankine" && unit != "celsius" && unit != "fahrenheit" && unit != "kelvin" {
		log.Fatalln("INVALID")
	}
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
	targetUnit, _ := unitReader.ReadString('\n')
	targetUnit = strings.ToLower(targetUnit)
	targetUnit = strings.TrimSuffix(targetUnit, "\n")

	//Allow teacher to input Student Resp
	respReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Student Response: ")
	inputResp, _ := respReader.ReadString('\n')
	inputResp = strings.TrimSuffix(inputResp, "\n")

	//Allow Input temperature parameters to be handled independently
	inputSplit := strings.Split(inputTemp, " ")

	//Set input Unit
	inputUnit := strings.ToLower(inputSplit[1])

	//Run unit validations
	unitValidation(inputUnit)
	unitValidation(targetUnit)

	//Set and error check input temp
	temperature := strToFloat(inputSplit[0])
	response := strToFloat(inputResp)

	//Selects which conversion to run based on input unit
	switch {
	case inputUnit == "kelvin":
		fmt.Println("The student is " + checkResp(kelvinConversions(temperature, targetUnit), response))
	case inputUnit == "celsius":
		fmt.Println("The student is " + checkResp(celsiusConversions(temperature, targetUnit), response))
	case inputUnit == "fahrenheit":
		fmt.Println("The student is " + checkResp(fahrenheitConversions(temperature, targetUnit), response))
	case inputUnit == "rankine":
		fmt.Println("The student is " + checkResp(rankineConversions(temperature, targetUnit), response))
	}
}
