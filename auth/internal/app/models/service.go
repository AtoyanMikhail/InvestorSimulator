package models

import "context"

type Service interface {
	SignUp(ctx context.Context, user *User) error
	LogIn(ctx context.Context, email, hashedPassword string) (string, string, error)
	Auhtenicate(ctx context.Context, token string) (string, error)
}
