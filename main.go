package main

import (
	//"log"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"goyave.dev/goyave/v4"
	"goyave.dev/goyave/v4/validation"
)

type User struct {
	Id       int    `json: "id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Age      int    `json: "age"`
	Email    string `json:"email"`
}

var users []User

func main() {
	/*if err := connectToMongo(); err != nil {
		log.Fatal(err)
	}
	if err := goyave.Start(routes); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}*/

	users = append(users, User{Id: 1, Name: "Jose", LastName: "Ibarria", Age: 38, Email: "topete@hotm.com"})
	users = append(users, User{Id: 2, Name: "Luz", LastName: "Vazquez", Age: 42, Email: "topete@hotm.com"})
	users = append(users, User{Id: 3, Name: "Bastian", LastName: "Ibarria", Age: 1, Email: "topete@hotm.com"})
	users = append(users, User{Id: 4, Name: "Rodrigo", LastName: "Curry", Age: 35, Email: "rodri@gmail.com"})
	users = append(users, User{Id: 5, Name: "James", LastName: "Fields", Age: 35, Email: "James_123@hotmail.com"})
	if err := goyave.Start(routes); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}

}

var PostRequest = validation.RuleSet{
	"Id":       validation.List{"required", "numeric"},
	"Name":     validation.List{"required", "string"},
	"LastName": validation.List{"required", "string"},
	"Age":      validation.List{"required", "numeric"},
	"Email":    validation.List{"required", "string"},
}

func routes(router *goyave.Router) {
	router.Post("/ApiRestGoyave", createUsers).Validate(PostRequest)
	router.Get("/ApiRestGoyave", readUsers)
	router.Get("/ApiRestGoyave/{id}", readUsers)
	router.Put("/ApiRestGoyave/{id}", updateUsers)
	router.Delete("/ApiRestGoyave/{id}", deleteUsers)

}

func createUsers(response *goyave.Response, request *goyave.Request) {
	var newUser User
	if err := request.ToStruct(&PostRequest); err != nil {
		response.JSON(http.StatusBadRequest, "Invalid user data")
		return
	}
	fmt.Println("users: ", users)
	users = append(users, newUser)
	response.JSON(http.StatusOK, "Successfully created user")
}

func readUsers(response *goyave.Response, request *goyave.Request) {
	id, ok := request.Params["id"]
	if !ok {
		response.JSON(http.StatusOK, users)
		return
	}
	idInt, _ := strconv.Atoi(id)
	var user User
	for _, u := range users {
		if u.Id == idInt {
			user = u
			break
		}
	}
	if user.Id == 0 {
		response.JSON(http.StatusNotFound, "User not found")
	} else {
		response.JSON(http.StatusOK, user)
	}
}

func updateUsers(response *goyave.Response, request *goyave.Request) {
	id, _ := strconv.Atoi(request.Params["id"])
	var updatedUser User
	if err := request.ToStruct(&updatedUser); err != nil {
		response.JSON(http.StatusBadRequest, "Invalid user data")
		return
	}
	for i, user := range users {
		if user.Id == id {
			users[i] = updatedUser
			response.JSON(http.StatusOK, "Suscelfull register")
			return
		}
	}
	response.JSON(http.StatusBadRequest, "The user doesn´t exist")
}

func deleteUsers(response *goyave.Response, request *goyave.Request) {
	id, _ := strconv.Atoi(request.Params["id"])
	for i, user := range users {
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)
			response.JSON(http.StatusOK, "Successfully deleted user")
			return
		}
	}
	response.JSON(http.StatusBadRequest, "The Register does not exist")
}
