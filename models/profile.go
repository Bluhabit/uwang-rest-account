package models

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
