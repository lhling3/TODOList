package service

import (
	"TODOList/store"
	"errors"
)

type UserService struct{}

func (s UserService) Create(user *store.User) (uint64, error) {
	if store.UserMap[user.Id] != nil {
		return 0, errors.New("该用户已经存在")
	}
	store.UserMap[user.Id] = user
	return user.Id, nil
}

func (s UserService) Delete(userId uint64) error {
	if store.UserMap[userId] == nil {
		return errors.New("用户不存在")
	}
	delete(store.UserMap, userId)
	return nil
}

func (s UserService) UpdatePassword(userId uint64, password string) error {
	user, ok := store.UserMap[userId]
	if !ok {
		return errors.New("用户不存在")
	}
	user.Password = password
	return nil
}
