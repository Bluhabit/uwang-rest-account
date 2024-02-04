package models

type UpdateProfileUsername struct {
	Username string `json:"username"`
}

type UpdateProfilePicture struct {
	ProfilePicture string `json:"profile_picture"`
}
