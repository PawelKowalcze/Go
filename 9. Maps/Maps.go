// ages := make(map[string]int) // create a map with keys of type string and values of type int
// ages["alice"] = 31
// ages["charlie"] = 34
// ages["charlie"] = 22

//diffAges = map[string]int{
//"John": 37,
//"Mary": 21,
//}

//// The len() function works on a map, it returns the total number of key/value pairs.
//ages = map[string]int{
//"John": 37,
//"Mary": 21,
//}
//fmt.Println(len(ages)) // 2

package main

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	maps := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("names and phoneNumbers have different length")
	}

	for i, name := range names {
		if name == "" {
			return nil, fmt.Errorf("name cannot be empty")
		}
		maps[name] = user{name: name, phoneNumber: phoneNumbers[i]}

	}
	return maps, nil
}

// don't touch below this line

type user struct {
	name        string
	phoneNumber int
}

func test(names []string, phoneNumbers []int) {
	fmt.Println("Creating map...")
	defer fmt.Println("====================================")
	users, err := getUserMap(names, phoneNumbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, name := range names {
		fmt.Printf("key: %v, value:\n", name)
		fmt.Println(" - name:", users[name].name)
		fmt.Println(" - number:", users[name].phoneNumber)
	}
}

func main() {
	test(
		[]string{"John", "Bob", "Jill"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	test(
		[]string{"John", "Bob"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	test(
		[]string{"George", "Sally", "Rich", "Sue"},
		[]int{20955559812, 38385550982, 48265554567, 16045559873},
	)
}
