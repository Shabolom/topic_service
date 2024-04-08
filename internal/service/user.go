package service

import (
	"github.com/gofrs/uuid"
	"service_topic/internal/domain"
	"service_topic/internal/models"
	"service_topic/internal/repository"
	"time"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

var userRepo = repository.NewUserRepo()

func (us UserService) Register(user models.User) error {
	id, _ := uuid.FromString(user.UserID)

	userEntity := domain.User{
		ID:          id,
		WhenCreated: time.Now(),
		UserName:    user.UserName,
	}

	err := userRepo.Register(userEntity)
	if err != nil {
		return err
	}

	return nil
}
