package main

import "fmt"

func main() {
	fmt.Println("starting Textio Server")

	messagesFromDoris := []string{
		"Hello",
		"Hi",
		"Good Morning",
		"Good Evening",
	}

	numMessages := float64(len(messagesFromDoris))
	costPerMessage := 0.2

	totalCost := numMessages * costPerMessage

	fmt.Printf("Total cost is %v \n", totalCost)

	var usernames string = "wagslane"
	var password string = "123456789"

	fmt.Println("Authorization: Basic", usernames+":"+password)

	//Variables
	/*
		bool
		string
		int  int8  int16  int32  int64
		uint uint8 uint16 uint32 uint64 uintptr
		byte // alias for uint8
		rune // alias for int32 -> like char in C
		float32 float64
		complex64 complex128
	*/
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	/*printf formats:
	%v - default representation
	%s - string
	%d - decimal
	%f - float,  %.2f - float with 2 decimal points
	*/
	fmt.Printf("%v %f %v %q \n",
		smsSendingLimit,
		costPerSMS,
		hasPermission,
		username)

	// short assignment operator:  :=

	congrats := "Congratulations"
	fmt.Println(congrats)

	penniesPerText := 2
	fmt.Printf("The type of penniesPerText is %T \n", penniesPerText) // %T is used to print the type of the variable

	averageOpenRate, displayMessage := 0.23, "is the average open rate"
	fmt.Println(averageOpenRate, displayMessage)

	accountAge := 2.6
	accountAgeInt := int(accountAge)
	fmt.Println(accountAgeInt)

	// cant change const variable

	const premiumPlanName string = "Premium"
	// premiumPlanName = "Gold" // this will throw an error

	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour
	fmt.Println(secondsInHour)

	// Sprintf - returns a formatted string

	const name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hello %s, your open rate is %.2f", name, openRate)
	fmt.Println(msg)

	// If statements

	messageLen := 30
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length: ", messageLen)

	if messageLen >= maxMessageLen {
		fmt.Println("Message too long")
	} else {
		fmt.Println("Message sent")
	}

	if lenght := len("Hello World"); lenght > 5 {
		fmt.Println("That's a long string")
	}

}
