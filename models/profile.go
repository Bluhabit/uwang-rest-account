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

type ListUserResponse struct {
	Id           string                `json:"id"`
	Email        string                `json:"email"`
	FullName     string                `json:"full_name"`
	UserName     string                `json:"username"`
	DateOfBirth  time.Time             `json:"dateOfBirth"`
	AuthProvider string                `json:"authProvider"`
	Status       string                `json:"status"`
	CreatedAt    time.Time             `json:"createdAt"`
	UpdatedAt    time.Time             `json:"updatedAt"`
	Deleted      bool                  `json:"deleted"`
}

type Pagination struct {
	Limit      int                `json:"limit,omitempty;query:limit"`
	Page       int                `json:"page,omitempty;query:page"`
	Sort       string             `json:"sort,omitempty;query:sort"`
	TotalRows  int64              `json:"total_rows"`
	TotalPages int                `json:"total_pages"`
	Rows       []ListUserResponse `json:"rows"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}