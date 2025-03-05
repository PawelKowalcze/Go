package main

import "fmt"

func concat(s1, s2 string) string {
	return s1 + s2
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	return sendsSoFar + sendsToAdd
}

func getNames() (string, string) {
	return "John", "Doe"
}

// I can name retur values - in order to document the return values

func getCoords() (x, y int) {
	//x, y are initialized to 0
	return // automatically returns x, y
}

func yearsUntilEvents(age int) (
	yearsUntilAdult int,
	yearsUntilDrinking int,
	untilDeath int,
) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 25 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	untilDeath = 100 - age
	if untilDeath < 0 {
		fmt.Println("You are already dead")
	}

	return yearsUntilAdult, yearsUntilDrinking, untilDeath
}

// Early returns

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return dividend / divisor, nil
}

func main() {
	fmt.Println(concat("starting Textio Server ", "Hello"))

	sendsSoFar := 430
	const sendsToAdd = 10
	fmt.Println("Total sends so far: ", sendsSoFar)

	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)

	fmt.Println("Total sends so far: ", sendsSoFar)

	// Ignoring return values

	firstName, _ := getNames() // _ - ignore the second return value
	fmt.Println("First Name: ", firstName)
	// fmt.Println("Last Name: ", lastName)

	yourAge := 101
	yearsUntilAdult, yearsUntilDrinking, untilDeath := yearsUntilEvents(yourAge)
	fmt.Println("Years until adult: ", yearsUntilAdult)
	fmt.Println("Years until drinking: ", yearsUntilDrinking)
	fmt.Println("Years until death: ", untilDeath)

	// Early returns
	x := 10.0
	y := 0.0

	result, err := divide(x, y)

	fmt.Println("Result: ", result)
	fmt.Println("Error: ", err)

}
