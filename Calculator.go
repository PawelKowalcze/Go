package main

import "fmt"

func main() {
	var numberOne float32
	var numberTwo float32
	var operationType int
	var programmEnd bool = true
	for {
		fmt.Println("Hi, this is your Calculator")
		fmt.Println("Please, choose your operation")
		fmt.Println("1. Addition \n2. Subtraction \n3. Multiplication \n4. Division")
		fmt.Scanf("%d", &operationType)
		switch operationType {
		case 1:
			fmt.Println("You have chosen Addition")
			fmt.Println("Select your numbers")
			fmt.Scanf("%f", &numberOne)
			fmt.Scanf("%f", &numberTwo)
			sum := numberOne + numberTwo
			fmt.Printf("The sum is %v \n", sum)
		case 2:
			fmt.Println("You have chosen Subtraction")
			fmt.Println("Select your numbers")
			fmt.Scanf("%f", &numberOne)
			fmt.Scanf("%f", &numberTwo)
			sum := numberOne - numberTwo
			fmt.Printf("The difference is %v \n", sum)
		case 3:
			fmt.Println("You have chosen Multiplication")
			fmt.Println("Select your numbers")
			fmt.Scanf("%f", &numberOne)
			fmt.Scanf("%f", &numberTwo)
			sum := numberOne * numberTwo
			fmt.Printf("The multiplication is %v \n", sum)
		case 4:
			fmt.Println("You have chosen Division")
			fmt.Println("Select your numbers")
			fmt.Scanf("%f", &numberOne)
			fmt.Scanf("%f", &numberTwo)
			sum := numberOne / numberTwo
			fmt.Printf("The division is %v \n", sum)
		default:
			fmt.Println("Please enter a valid operation type.")
		}
		fmt.Println("Do you want to continue?\n 1.Yes 0.No")
		fmt.Scanf("%t\n", &programmEnd)
		if programmEnd == false {
			fmt.Println("Programm has ended")
			return
		}
	}
}
