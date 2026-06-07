package user

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

var users = []User{
	{
		Id:    1,
		Name:  "John Doe",
		Email: "jD0Hw@example.com",
		Age:   30,
	},
	{
		Id:    2,
		Name:  "Jane Doe",
		Email: "2b4e9@example.com",
		Age:   25,
	},
	{
		Id:    3,
		Name:  "Bob Smith",
		Email: "r9B2o@example.com",
		Age:   35,
	},
}
