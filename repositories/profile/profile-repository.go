package profile

import (
	"context"
	"fmt"
	"github.com/Bluhabit/uwang-rest-account/common"
	"github.com/Bluhabit/uwang-rest-account/entity"
	"github.com/Bluhabit/uwang-rest-account/models"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
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

func (repo *ProfileRespository) UpdateProfileUsername(sessionId string, username string) models.BaseResponse[string] {
	//prepare data for update or inserted
	var userCredential entity.UserCredential
	var response = models.BaseResponse[string]{}

	//ambil detail user dari redis
	redis_key := common.CreateRedisKeyUserSession(sessionId)
	session := repo.cache.HGetAll(context.Background(), redis_key)
	if session == nil {
		return response.BadRequest("", "Sesi tidak ditemukan")
	}
	user := session.Val()
	userId := user["user_id"]
	if len(userId) < 1 {
		return response.BadRequest("", "Sesi tidak ditemukan[1]")
	}

	//cek apakah usernya ada di database
	userError := repo.db.Where("id = ?", userId).First(&userCredential).Error
	if userError != nil {
		return response.BadRequest("", "Akun tidak ditemukan")
	}

	//jika error tidak kosong berarti ada yang pakai
	var existingUser entity.UserCredential
	usernameError := repo.db.Where("username=?", username).First(&existingUser).Error
	if usernameError == nil {
		return response.BadRequest("", "Username tidak dapat digunakan.")
	}

	//update username ke database
	userCredential.Username = username
	repo.db.Save(userCredential)
	return response.Success("", "Berhasil merubah username.")
}

func (repo *ProfileRespository) UpdateProfilePicture(sessionId string, profilePicture string) models.BaseResponse[string] {
	//prepare data for update or inserted
	var userProfile entity.UserProfile
	var response = models.BaseResponse[string]{}

	//ambil userId dari redis
	redisKey := common.CreateRedisKeyUserSession(sessionId)
	session := repo.cache.HGetAll(context.Background(), redisKey)
	if session == nil {
		return response.BadRequest("", "Sesi tidak ditemukan")
	}
	user := session.Val()
	userId := user["user_id"]
	if len(userId) < 1 {
		return response.BadRequest("", "Sesi tidak ditemukan[1]")
	}

	fmt.Printf("User Id %s", userId)
	//jika belum ada buat data baru
	if err := repo.db.Where("user_id = ? AND key = 'profile-picture'", userId).First(&userProfile).Error; err != nil {
		var profilePictureID = uuid.NewString()
		userProfile := entity.UserProfile{
			ID:        profilePictureID,
			Key:       "profile-picture",
			Value:     profilePicture,
			UserID:    userId,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Deleted:   false,
		}
		repo.db.Save(userProfile)
		return response.Success("", "Berhasil merubah foto profil.")
	}
	//update data lama
	userProfile.Value = profilePicture
	repo.db.Save(userProfile)
	return response.Success("", "Berhasil merubah foto profil.")
}

func (repo *ProfileRespository) UpdateProfileTopics(sessionId string, topics string) models.BaseResponse[string] {
	//prepare data for update or inserted
	var userProfile entity.UserProfile
	var response = models.BaseResponse[string]{}

	//ambil userId dari redis
	redis_key := common.CreateRedisKeyUserSession(sessionId)
	session := repo.cache.HGetAll(context.Background(), redis_key)
	if session == nil {
		return response.BadRequest("", "Sesi tidak ditemukan")
	}
	user := session.Val()
	userId := user["user_id"]
	if len(userId) < 1 {
		return response.BadRequest("", "Sesi tidak ditemukan[1]")
	}

	//jika belum ada buat data baru
	if err := repo.db.Where("user_id = ? AND key = 'topics'", userId).First(&userProfile).Error; err != nil {
		var profilePictureID = uuid.NewString()
		userProfile := entity.UserProfile{
			ID:        profilePictureID,
			Key:       "topics",
			Value:     topics,
			UserID:    userId,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Deleted:   false,
		}
		repo.db.Save(userProfile)
		return response.Success("", "Berhasil merubah foto profil.")
	}

	//update data lama
	userProfile.Value = topics
	repo.db.Save(userProfile)
	return response.Success("", "Berhasil membuat topik profil")
}

func (repo *ProfileRespository) UpdateProfileLevel(sessionId string, level string) models.BaseResponse[string] {
	//prepare data for update or inserted
	var userProfile entity.UserProfile
	var response = models.BaseResponse[string]{}

	//ambil userId dari redis
	redis_key := common.CreateRedisKeyUserSession(sessionId)
	session := repo.cache.HGetAll(context.Background(), redis_key)
	if session == nil {
		return response.BadRequest("", "Sesi tidak ditemukan")
	}
	user := session.Val()
	userId := user["user_id"]
	if len(userId) < 1 {
		return response.BadRequest("", "Sesi tidak ditemukan[1]")
	}

	//jika belum ada buat data baru
	if err := repo.db.Where("user_id = ? AND key = 'level'", userId).First(&userProfile).Error; err != nil {
		var profilePictureID = uuid.NewString()
		userProfile := entity.UserProfile{
			ID:        profilePictureID,
			Key:       "level",
			Value:     level,
			UserID:    userId,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Deleted:   false,
		}
		repo.db.Save(userProfile)
		return response.Success("", "Berhasil menyimpan topic")
	}

	//update data lama
	userProfile.Value = level
	repo.db.Save(userProfile)
	return response.Success("", "Berhasil menyimpan topic")
}
