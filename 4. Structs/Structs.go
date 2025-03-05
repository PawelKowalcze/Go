package main

import "fmt"

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

// anonymous struct - struct without a name
type car struct {
	Make   string
	Model  string
	Height int
	Width  int
	//We can also have a struct inside a struct
	Wheel struct {
		Radius   int
		Material string
	}
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.number > 0 && mToSend.sender.name != "" && mToSend.recipient.number > 0 && mToSend.recipient.name != "" {
		return true
	}
	return false
}

// Embedded struct - struct inside a struct

type truck struct {
	car     //embedded struct - it's only a type - we inherit all the fields and methods of the car struct
	weight  int
	bedSize int
}

func main() {
	// Nested struct
	lanesTruck := truck{
		bedSize: 10,
		car: car{
			Make:  "toyota",
			Model: "camry",
		},
	}

	fmt.Println(lanesTruck.bedSize)

	// embedded fields promoted to the top-level
	// instead of lanesTruck.car.make
	fmt.Println(lanesTruck.Make)
	fmt.Println(lanesTruck.Model)

	myCar := struct {
		make  string
		model string
	}{
		make:  "Ford",
		model: "F150",
	}

	fmt.Println(myCar)
	fmt.Println(canSendMessage(messageToSend{
		message: "Hello",
		sender: user{
			name:   "John",
			number: 1234,
		},
		recipient: user{
			name:   "Doe",
			number: 987654321,
		},
	}))
}
