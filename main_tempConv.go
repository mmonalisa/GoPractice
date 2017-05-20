package main

import "fmt"

func main() {
	fmt.Println("Enter the temperature in Fahraneit:")
	var input float64
	fmt.Scanf("%f", &input)
	output := (input - 32)*5/9
	fmt.Println(output)

}