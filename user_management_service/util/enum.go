package util

type UserType uint64
type Role uint64
type Permission uint64

const (
	Registration Role = 1 + iota
	Hostel
	Canteen
)

const (
	SuperAdmin UserType = 1 + iota
	InstitutionAdmin
	CollegeAdmin
)

const (
	Admin Permission = 1 + iota
	Moderator
	Viewer
)
