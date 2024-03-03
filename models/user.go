package models

import "time"

type UserCredentialResponse struct {
	Id           string    `json:"id"`
	Email        string    `json:"email"`
	FullName     string    `json:"full_name"`
	UserName     string    `json:"username"`
	DateOfBirth  time.Time `json:"dateOfBirth"`
	AuthProvider string    `json:"authProvider"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Deleted      bool      `json:"deleted"`
}