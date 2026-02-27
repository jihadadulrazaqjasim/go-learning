package models

type UserID int

type User struct {
	ID    UserID
	Name  string
	Email string
	Age   int
}

func (u User) IsAdult() bool {
	return u.Age >=18
}