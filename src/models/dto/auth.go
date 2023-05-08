package dto

import (
	"timetracker/models"
)

type ReqUserSignIn struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ReqUserSignUp struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	About    string `json:"about"`
	AvatarID uint64 `json:"avatar_id"`
	Password string `json:"password" validate:"required"`
}

func (req *ReqUserSignIn) ToModelUser() *models.User {
	return &models.User{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *ReqUserSignUp) ToModelUser() *models.User {
	return &models.User{
		Name:     req.Name,
		Email:    req.Email,
		About:    req.About,
		AvatarID: req.AvatarID,
		Password: req.Password,
	}
}

type RespUser struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	About    string `json:"about"`
	AvatarID uint64 `json:"avatar_id"`
}

func GetResponseFromModelUser(user *models.User) *RespUser {
	return &RespUser{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		About:    user.About,
		AvatarID: user.AvatarID,
	}
}

//
//func GetResponseFromModelEntries(entries []*models.Entry) []*RespEntry {
//	result := make([]*RespEntry, 0, 10)
//	for _, entry := range entries {
//		result = append(result, GetResponseFromModelEntry(entry))
//	}
//
//	return result
//}
