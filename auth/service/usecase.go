package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"myProject/auth"
	"myProject/models"
	"time"
)

const (
	salt       = "bcb545454yh5p5HG"
	tokenTTL   = 1 * time.Minute
	refreshTTL = 24 * 7 * time.Hour
	signingKey = "QHhpZGlF2DG3SD3F3G2SDF3H4vCg=="
	refreshKey = "ds7B989umHJ98opi;m2"
)

type AuthUsecase struct {
	repo auth.UserRepository
}

func NewAuthUsecase(repo auth.UserRepository) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}

func (a *AuthUsecase) SignUp(ctx context.Context, user models.User) (uint16, error) {
	user.Password = GeneratePasswordHash(user.Password)
	if err := isEmailValid(user.Email); err != nil {
		return 0, err
	}
	return a.repo.CreateUser(ctx, user)
}

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
