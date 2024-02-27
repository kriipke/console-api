package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        					uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      					string    `gorm:"type:varchar(255);not null"`
	Email     					string    `gorm:"uniqueIndex;not null"`
	Password  					string    `gorm:"not null"`
	Role      					string    `gorm:"type:varchar(255);not null"`
	Provider  					string    `gorm:"not null"`
	Photo     					string    `gorm:"not null"`
	VerificationCode 		string
	PasswordResetToken 	string
	PasswordResetAt    	time.Time
	Verified  					bool     	`gorm:"not null"`
	CreatedAt 					time.Time	`gorm:"not null"`
	UpdatedAt 					time.Time	`gorm:"not null"`
}

type SignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
	Photo           string `json:"photo" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// The User struct represents the SQL table in the database. When the user makes a password reset request to the Golang API, the server will generate the reset token, and store the hashed code in the PasswordResetToken column in the database before sending the unhashed code to the user’s email address.
//  For security reasons, the password reset code will have an expiry time of 15 minutes. Within this period, the user is expected to complete the password reset process else a new reset link has to be requested since the old one will be deleted from the database.

// ? ForgotPasswordInput struct
type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

// ? ResetPasswordInput struct
type ResetPasswordInput struct {
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

