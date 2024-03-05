package models

type UpdateProfileRequest struct {
	Username      string `json:"username" binding:"required"`
	FullName      string `json:"full_name" binding:"required"`
	Bio           string `json:"bio" binding:"required"`
	Link          string `json:"link" binding:"required"`
	InterestTopic string `json:"interest_topic" binding:"required"`
}

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

type UserProfileResponse struct {
	Id    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}
