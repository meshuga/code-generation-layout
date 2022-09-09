package model

//go:generate stringer -type=UserStatus
type UserStatus int

const (
	Unverified UserStatus = iota
	Verified
	Suspended
)

type UserResponse struct {
	ID     int64
	Email  string
	Status UserStatus
}

type UsersCollection []UserResponse
