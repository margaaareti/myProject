package auth

import (
	"context"
	"myProject/models"
)

const CtxUserId = "userId"
const CtxUserName = "name"
const CtxUserSurname = "surname"
const CtxUserPatronymic = "patronymic"

type UseCase interface {
	SignUp(ctx context.Context, user models.User) (uint16, error)
	//SignIn(ctx context.Context, name, password string) (*models.User, error)
}
