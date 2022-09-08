package model

//go:generate stringer -type=UserStatus

type UserStatus int

const (
	Unverified UserStatus = iota
	Verified
	Suspended
)
