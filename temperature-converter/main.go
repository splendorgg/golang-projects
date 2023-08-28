package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func celciustoFahrenheit(temp float64) {
	tempf := (temp * 9 / 5) + 32
	fmt.Printf("%.2f°C is %.2f°F\n", temp, tempf)
	fmt.Println("")
}
func fahrenheittoCelcius(temp float64) {
	tempc := (temp - 32) * 5 / 9
	fmt.Printf("%.2f°C is %.2f°F\n", temp, tempc)
	fmt.Println("")
}

func main() {
	fmt.Println("This is a temperature converter between Celcius and Fahrenheit ")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("1 for Celcius to Fahrenheit")
		fmt.Println("2 for Fahrenheit to Celcius")
		fmt.Println("3 for exit the program")
		fmt.Print("Enter your choice: ")
		scanner.Scan()
		temp := scanner.Text()

		switch temp {
		case "1":
			fmt.Print("Enter your Celcius value: ")
			scanner.Scan()
			tempvalue := scanner.Text()
			newtemp, _ := strconv.ParseFloat(tempvalue, 64)
			celciustoFahrenheit(newtemp)
		case "2":
			fmt.Print("Enter your Fahrenheit value: ")
			scanner.Scan()
			tempvalue := scanner.Text()
			newtemp, _ := strconv.ParseFloat(tempvalue, 64)
			fahrenheittoCelcius(newtemp)
		case "3":
			return
		default:
			fmt.Println("Invalid input")
		}
	}
}
