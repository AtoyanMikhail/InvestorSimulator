package models

// User представляет собой модель пользователя.
type User struct {
	ID           int64  `db:"id" json:"id"`
	Username     string `db:"username" json:"username" validate:"required,min=5,max=20"`
	Email        string `db:"email" json:"email" validate:"required,email, max=100"`
	PasswordHash string `db:"password_hash" json:"-" validate:"required"`
	RefreshToken string `db:"refresh_token" json:"refresh_token"`
}
