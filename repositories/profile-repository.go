package repositories

import (
	"github.com/Bluhabit/uwang-rest-account/common"
	"github.com/Bluhabit/uwang-rest-account/entity"
)

func GetUserProfileById(id string) *entity.UserCredential {
	db := common.GetDbConnection()
	if db == nil {
		return nil
	}

	var userCredential entity.UserCredential
	if err := db.Where("id = ?", id).
		First(&userCredential).Error; err != nil {
		return nil
	}
	return &userCredential
}

func UpdateUsername(userCredential *entity.UserCredential) error {
	db := common.GetDbConnection()
	return db.Save(userCredential).Error
}

func GetUserByUsername(username string) *entity.UserCredential {
	db := common.GetDbConnection()
	var user entity.UserCredential

	if err := db.Where("username = ?", username).
		First(&user).Error; err != nil {
		return nil
	}
	return &user
}

func GetProfilePictureByUserID(userId string) *entity.UserProfile {
	db := common.GetDbConnection()
	if db == nil {
		return nil
	}

	var UserProfile entity.UserProfile
	if err := db.Where("user_id = ? AND key = 'profile-picture'", userId).
		First(&UserProfile).Error; err != nil {
		return nil
	}
	return &UserProfile
}

func CreateUserProfile(userProfile *entity.UserProfile) error {
	db := common.GetDbConnection()

	return db.Create(userProfile).Error
}

func UpdateUserProfile(userProfile *entity.UserProfile) error {
	db := common.GetDbConnection()

	return db.Save(userProfile).Error
}

func GetProfileInterestTopicByUserID(userId string) *entity.UserProfile {
	db := common.GetDbConnection()
	if db == nil {
		return nil
	}

	var UserProfile entity.UserProfile
	if err := db.Where("user_id = ? AND key = 'interest-topic'", userId).
		First(&UserProfile).Error; err != nil {
		return nil
	}
	return &UserProfile
}
