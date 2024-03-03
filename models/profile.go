package models

import "time"

type UpdateProfileUsername struct {
	Username string `json:"username" binding:"required"`
}

type UpdateProfilePicture struct {
	ProfilePicture string `json:"url" binding:"required"`
}

type UpdateProfileInterestTopic struct {
	InterestTopic string `json:"topics" binding:"required"`
}

type UpdateProfileLevel struct {
	Level string `json:"level" binding:"required"`
}

type DetailUserResponse struct {
	Id           string                `json:"id"`
	Email        string                `json:"email"`
	Password     string                `json:"password"`
	FullName     string                `json:"full_name"`
	UserName     string                `json:"username"`
	DateOfBirth  time.Time             `json:"dateOfBirth"`
	AuthProvider string                `json:"authProvider"`
	Status       string                `json:"status"`
	CreatedAt    time.Time             `json:"createdAt"`
	UpdatedAt    time.Time             `json:"updatedAt"`
	Deleted      bool                  `json:"deleted"`
	UserProfile  []UserProfileResponse `json:"user_profile"`
}

type UserProfileResponse struct {
	Id    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

