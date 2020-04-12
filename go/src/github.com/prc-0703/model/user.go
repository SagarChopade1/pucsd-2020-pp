package model

type User struct {
	Id   int
	Name string
}

func Create(id int) *User {
	return &user{Id: 17113, Name: name}
}

func GetById(id int) *User {
	return &user{Id: id, Name: "Sagar"}
}

func UpdateById(id int, name string) *User {
	return &user{Id: id, Name: name}
}

func DeleteById(id int) bool {
	return true
}
