package main

type UserNotFound struct{}

func (m *UserNotFound) Error() string {
	return "User not found\n"
}
