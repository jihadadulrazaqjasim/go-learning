package main

import (
	"fmt"
	"usermgmt/models"
	"usermgmt/service"
)

func main() {

	// colors := make(map[string]string, 1000)
	// colors["red"] = "#ff0000"
	// colors["somecolor"] = "#ffbbbb"
	// colors := map[string]string { "white" : "#ffffff", "red" : "#ff0000"}
	// fmt.Println(colors,len(colors))


	userOne := models.User{ID: 1, Name: "Jihad", Email: "jihad@gmail.com", Age: 25}
	userTwo := models.User{ID: 2, Name: "Mohammad", Email: "moha@gmail.com", Age: 24}
	userThree := models.User{ID: 3, Name: "Yusf", Email: "yussa@gmail.com", Age: 15}

	userService2 := service.NewUserService()
	fmt.Println(userService2)

	userService2.AddUser(userOne)
	userService2.AddUser(userTwo)
	userService2.PrintAllUsers()

	//add 3
	userService := service.UserService{}
	userService.AddUser(userOne)
	userService.AddUser(userTwo)
	userService.AddUser(userThree)

	//fetch by id
	fetchedUser,err := userService.GetUserByID(1)
	if err == nil {
		fmt.Printf("Name: %v, Email: %+v", fetchedUser.Name, fetchedUser.Email)
	}

	//update
	userService.UpdateUserName(1, "Gynann")

	//delete
	userService.DeleteUser(2)
	fmt.Print(userService.DeleteUser(22))

	//print all
	userService.PrintAllUsers()
}
