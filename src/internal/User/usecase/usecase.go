package usecase

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	userRep "timetracker/internal/User/repository"
	"timetracker/models"
)

type UsecaseI interface {
	UpdateUser(e *models.User) error
	GetUser(id uint64) (*models.User, error)
}

type usecase struct {
	userRepository userRep.RepositoryI
}

func (u *usecase) UpdateUser(user *models.User) error {
	_, err := u.userRepository.GetUser(user.ID)
	if err != nil {
		return errors.Wrap(err, "user repository error")
	}

	if user.Password != "" { // TODO может быть лучше когда будет DTO из dto в models конвертить и сразу хешировать, чтобы в бизнеслогике этого не было
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
		if err != nil {
			return errors.Wrap(err, "Error in func user.Usecase.UpdateUser bcrypt")
		}

		user.Password = string(hashedPassword)
	}

	err = u.userRepository.UpdateUser(user)
	if err != nil {
		return errors.Wrap(err, "Error in func user.Usecase.UpdateUser")
	}

	return nil
}

func (u *usecase) GetUser(id uint64) (*models.User, error) {
	user, err := u.userRepository.GetUser(id)
	if err != nil {
		return nil, errors.Wrap(err, "Error in func user.Usecase.GetUser")
	}
	user.Password = ""

	return user, nil
}

func New(uRep userRep.RepositoryI) UsecaseI {
	return &usecase{
		userRepository: uRep,
	}
}
