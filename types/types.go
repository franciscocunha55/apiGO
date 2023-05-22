package types

import "math/rand"

type User struct {
	Id   int
	Name string
	age  int
}

func NewUser(name string, age int) *User {
	return &User{
		Id:   rand.Intn(10),
		Name: name,
		age:  age,
	}
}

func ChangeName(user *User, newName string) {
	user.Name = newName
}
