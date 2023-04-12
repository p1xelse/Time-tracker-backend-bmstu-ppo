package models

type Tag struct {
	ID     uint64 `json:"id"`
	UserID uint64 `json:"user_id"`
	Name   string `json:"name"`
	About  string `json:"about"`
	Color  string `json:"color"`
}
