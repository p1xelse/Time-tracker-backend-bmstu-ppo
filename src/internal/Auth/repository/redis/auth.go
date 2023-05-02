package redis

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"timetracker/internal/Auth/repository"
	"timetracker/models"
)

type authRepository struct {
	db *redis.Client
}

func (ar authRepository) CreateCookie(cookie *models.Cookie) error {
	err := ar.db.Set(cookie.SessionToken, cookie.UserID, cookie.MaxAge).Err()

	if err != nil {
		return errors.Wrap(err, "redis error")
	}

	return nil
}

func (ar authRepository) GetUserByCookie(value string) (string, error) {
	userIdStr, err := ar.db.Get(value).Result()

	if errors.Is(err, redis.Nil) {
		return "", models.ErrNotFound
	} else if err != nil {
		return "", errors.Wrap(err, "redis error")
	}

	return userIdStr, nil
}

func (ar authRepository) DeleteCookie(value string) error {
	err := ar.db.Del(value).Err()

	if err != nil {
		return errors.Wrap(err, "redis error")
	}

	return nil
}

func NewAuthRepository(db *redis.Client) repository.RepositoryI {
	return &authRepository{
		db: db,
	}
}
