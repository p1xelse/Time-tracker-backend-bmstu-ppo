package models

type Cookie struct {
	SessionToken string
	UserID       uint64
	MaxAge       int64
}
