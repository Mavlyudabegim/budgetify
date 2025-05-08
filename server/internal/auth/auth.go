package auth

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"server/internal/auth/helper"

	"server/internal/database"
)

type AuthService struct {
	Q *database.Queries
}

func (a *AuthService) RegisterUser(ctx context.Context, email, password string) (*database.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user, err := a.Q.CreateUser(ctx, database.CreateUserParams{Email: email, Password: string(hashed)})

	if err != nil {
		if helper.IsUniqueViolation(err) {
			return nil, fmt.Errorf("email already exists")
		}

		log.Printf("RegisterUser error: %+v", err) // <- ADD THIS
		return nil, err
	}

	return &user, nil
}

func (a *AuthService) LoginUser(ctx context.Context, email, password string) (*database.User, error) {
	user, err := a.Q.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return &user, nil
}
