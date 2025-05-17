package user_service

import "blogW_server/models"

type UserService struct {
	userModel models.UserModel
}

func NewUserService(user models.UserModel) *UserService {
	return &UserService{
		userModel: user,
	}
}
