package repository

import (
	"context"
	"fmt"
	"service_topic/config"
	"service_topic/internal/domain"
)

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) Register(user domain.User) error {
	sql, args, err := config.Sq.
		Insert("users").
		Columns("id", "when_created", "user_name").
		Values(user.ID, user.WhenCreated, user.UserName).
		ToSql()
	if err != nil {
		return err
	}
	fmt.Println(sql, args)
	_, err = config.Pool.Exec(context.TODO(), sql, args...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
