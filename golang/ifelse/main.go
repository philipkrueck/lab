package main

import "fmt"

func main() {
	weight := 83.1 // kg
	height := 1.79 // meters

	if bmi := weight / (height * height); bmi < 18.5 {
		fmt.Printf("BMI: %.1f - Underweight\n", bmi)
	} else if bmi < 26 {
		fmt.Printf("BMI: %.1f - Normal weight\n", bmi)
	} else if bmi < 30 {
		fmt.Printf("BMI: %.1f - Overweight\n", bmi)
	} else {
		fmt.Printf("BMI: %.1f - Obese\n", bmi)
	}
}
