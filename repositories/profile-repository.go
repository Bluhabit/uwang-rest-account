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

func UpdateUserProfile(userCredential *entity.UserCredential) error {
	db := common.GetDbConnection()
	return db.Save(userCredential).Error
}
