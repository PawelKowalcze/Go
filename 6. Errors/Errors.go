package main

import "fmt"

type user struct {
	ID int
}

func getUser() (user, error) {
	//get user logic
	return user{}, nil
}

func getProfilePic(userID int) (string, error) {
	return fmt.Sprint(userID), nil
}

func main() {
	user, err := getUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	profilePic, err := getProfilePic(user.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(profilePic)
}
