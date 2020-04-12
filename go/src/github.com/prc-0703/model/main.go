package main

import (
	"fmt"

	"github.com/prc-0703/model"
)

func main() {
	fmt.Println("Hello world")
	//user := &model.User{Name:"sagar",ID:17114} //on heap and send parameter
	//user := new(model.User) //on heap
	var user model.User //loacl
	fmt.Println("id", user.Id)
	fmt.Println("name", user.Name)
}
