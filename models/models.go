package models

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

var users = []User{
	{
		Id:    1,
		Name:  "Alice",
		Email: "alice@example.com",
		Age:   30,
	},
	{
		Id:    2,
		Name:  "Bob",
		Email: "bob@example.com",
		Age:   23,
	},
}
