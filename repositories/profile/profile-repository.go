package profile

import (
	"context"

	"github.com/Bluhabit/uwang-rest-account/common"
	"github.com/Bluhabit/uwang-rest-account/entity"
	"github.com/Bluhabit/uwang-rest-account/models"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ProfileRespository struct {
	db    *gorm.DB
	cache *redis.Client
	minio *minio.Client
}

func Init() *ProfileRespository {
	dbConn := common.GetDbConnection()
	redis := common.GetRedisConnection()
	minio := common.GetMinioClient()

	return &ProfileRespository{
		db:    dbConn,
		cache: redis,
		minio: minio,
	}

}
func (repo *ProfileRespository) GetUserProfileById(id string) *entity.UserCredential {
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

func (repo *ProfileRespository) UpdateUsername(userCredential *entity.UserCredential) error {
	db := common.GetDbConnection()
	return db.Save(userCredential).Error
}

func (repo *ProfileRespository) GetUserByUsername(username string) *entity.UserCredential {
	db := common.GetDbConnection()
	var user entity.UserCredential

	if err := db.Where("username = ?", username).
		First(&user).Error; err != nil {
		return nil
	}
	return &user
}

func (repo *ProfileRespository) UpdateProfileUsername(session_id string, username string) models.BaseResponse[string] {
	var UserProfile entity.UserProfile
	var response = models.BaseResponse[string]{}

	redis_key := common.CreateRedisKeyUserSession(session_id)
	session := repo.cache.HGetAll(context.Background(), redis_key)
	if session == nil {
		return response.BadRequest("", "Sesi tidak ditemukan")
	}
	user := session.Val()
	userId := user["user_id"]
	if len(userId) < 1 {

	}

	if err := repo.db.Where("user_id = ? AND key = 'profile-picture'", userId).
		First(&UserProfile).Error; err != nil {
		return response.BadRequest("", "User tidak ditemukan.")
	}

}

func (repo *ProfileRespository) CreateUserProfile(userProfile *entity.UserProfile) error {
	db := common.GetDbConnection()

	return db.Create(userProfile).Error
}

func (repo *ProfileRespository) UpdateUserProfile(userProfile *entity.UserProfile) error {
	db := common.GetDbConnection()

	return db.Save(userProfile).Error
}

func (repo *ProfileRespository) GetProfileInterestTopicByUserID(userId string) *entity.UserProfile {
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

func (repo *ProfileRespository) GetProfileLevelByUserID(userId string) *entity.UserProfile {
	db := common.GetDbConnection()
	if db == nil {
		return nil
	}

	var UserProfile entity.UserProfile
	if err := db.Where("user_id = ? AND key = 'level'", userId).
		First(&UserProfile).Error; err != nil {
		return nil
	}
	return &UserProfile
}
