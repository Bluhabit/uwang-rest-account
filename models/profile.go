package models

type UpdateProfileUsername struct {
	Username string `json:"username"`
}

type UpdateProfilePicture struct {
	ProfilePicture string `json:"profile_picture"`
}

type UpdateProfileInterestTopic struct {
	InterestTopic string `json:"interest_topic"`
}
