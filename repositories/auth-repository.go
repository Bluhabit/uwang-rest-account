package repositories

import (
	"github.com/Bluhabit/uwang-rest-account/common"
	"github.com/Bluhabit/uwang-rest-account/entity"
)

func GetUserById(userId string) *entity.UserCredential {
	db := common.GetDbConnection()
	var user entity.UserCredential
	data := db.First(&user, userId)

	if data.Error != nil {

	}
	return &user
}
