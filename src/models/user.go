package models

type User struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	About    string `json:"about"`
	AvatarID uint64 `json:"avatar_id"`
	Password string `json:"password"`
}
