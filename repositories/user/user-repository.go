package user

import (
	"strconv"

	"github.com/Bluhabit/uwang-rest-account/common"
	"github.com/Bluhabit/uwang-rest-account/entity"
	"github.com/Bluhabit/uwang-rest-account/models"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserRespository struct {
	db    *gorm.DB
	cache *redis.Client
	minio *minio.Client
}

func Init() *UserRespository {
	dbConn := common.GetDbConnection()
	redis := common.GetRedisConnection()
	minio := common.GetMinioClient()

	return &UserRespository{
		db:    dbConn,
		cache: redis,
		minio: minio,
	}
}

func Paginate(page string, size string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(page)
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(size)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (repo *UserRespository) GetListUser(page string, size string) models.BaseResponseWithPagination[[]models.UserCredentialResponse] {
	var userCredential []entity.UserCredential
	var UserCredentialResponse []models.UserCredentialResponse
	var response = models.BaseResponseWithPagination[[]models.UserCredentialResponse]{}

	err := repo.db.Scopes(Paginate(page, size)).Find(&userCredential).Error
	if err != nil {
		return response.BadRequestPagination(UserCredentialResponse, "Tidak dapat mengambil data",page,size,0)
	}
	for _, userCredential := range userCredential {
		UserCredentialResponse = append(UserCredentialResponse, models.UserCredentialResponse{
			Id:           userCredential.ID,
			Email:        userCredential.Email,
			FullName:     userCredential.FullName,
			UserName:     userCredential.Username,
			DateOfBirth:  userCredential.DateOfBirth,
			AuthProvider: userCredential.AuthProvider,
			Status:       userCredential.Status,
			CreatedAt:    userCredential.CreatedAt,
			UpdatedAt:    userCredential.UpdatedAt,
			Deleted:      userCredential.Deleted,
		})

	}
	return response.SuccessWithPagination(
		UserCredentialResponse, 
		"Berhasil mengambil data",
		page,
		size,
		len(userCredential),
	)

}

func (repo *UserRespository) GetListUserByQuery(query string) models.BaseResponse[[]models.UserCredentialResponse] {
	var userCredential []entity.UserCredential
	var userCredentialResponse []models.UserCredentialResponse
	var response = models.BaseResponse[[]models.UserCredentialResponse]{}

	err := repo.db.Where("full_name LIKE ?", "%"+query+"%").Find(&userCredential).Error
	if err != nil {
		return response.BadRequest(userCredentialResponse, "Tidak dapat menemukan data")
	}
	if len(userCredential) == 0 {
		return response.BadRequest(userCredentialResponse, "User tidak ditemukan")
	}

	for _, userCredential := range userCredential {
		userCredentialResponse = append(userCredentialResponse, models.UserCredentialResponse{
			Id:           userCredential.ID,
			Email:        userCredential.Email,
			FullName:     userCredential.FullName,
			DateOfBirth:  userCredential.DateOfBirth,
			AuthProvider: userCredential.AuthProvider,
			Status:       userCredential.Status,
			CreatedAt:    userCredential.CreatedAt,
			UpdatedAt:    userCredential.UpdatedAt,
			Deleted:      userCredential.Deleted,
		})
	}
	return response.Success(userCredentialResponse, "Berhasil menemukan data")
}

func (repo *UserRespository) GetDetailUser(userId string) models.BaseResponse[models.UserCredentialResponse] {
	//prepare data
	var userCredential entity.UserCredential
	var userProfile []entity.UserProfile
	var responseDetailUser = models.UserCredentialResponse{}
	var response = models.BaseResponse[models.UserCredentialResponse]{}

	if err := repo.db.Where("id = ?", userId).First(&userCredential).Error; err != nil {
		return response.BadRequest(responseDetailUser, "Sesi tidak ditemukan")
	}
	if err := repo.db.Where("id = ?", userId).Find(&userProfile).Error; err != nil {
		return response.BadRequest(responseDetailUser, "Sesi tidak ditemukan")
	}

	var userProfileResponse []models.UserProfileResponse

	for _, profile := range userProfile {
		userProfileResponse = append(userProfileResponse, models.UserProfileResponse{
			Id:    profile.ID,
			Key:   profile.Key,
			Value: profile.Value,
		})
	}

	responseDetailUser = models.UserCredentialResponse{
		Id:           userCredential.ID,
		Email:        userCredential.Email,
		FullName:     userCredential.FullName,
		UserName:     userCredential.Username,
		DateOfBirth:  userCredential.DateOfBirth,
		AuthProvider: userCredential.AuthProvider,
		Status:       userCredential.Status,
		CreatedAt:    userCredential.CreatedAt,
		UpdatedAt:    userCredential.UpdatedAt,
		Deleted:      userCredential.Deleted,
		UserProfile:  userProfileResponse,
	}

	return response.Success(responseDetailUser, "Berhasil mengambil detail user")
}

func (repo *UserRespository) GetTopUser() models.BaseResponse[[]models.UserCredentialResponse] {
	var userCredential []entity.UserCredential
	var UserCredentialResponse []models.UserCredentialResponse
	var response = models.BaseResponse[[]models.UserCredentialResponse]{}

	err := repo.db.Where("ORDER By created_at DESC").Find(&userCredential).Error
	if err != nil {
		return response.BadRequest(UserCredentialResponse, "Tidak dapat mengambil data")
	}
	for _, userCredential := range userCredential {
		UserCredentialResponse = append(UserCredentialResponse, models.UserCredentialResponse{
			Id:           userCredential.ID,
			Email:        userCredential.Email,
			FullName:     userCredential.FullName,
			UserName:     userCredential.Username,
			DateOfBirth:  userCredential.DateOfBirth,
			AuthProvider: userCredential.AuthProvider,
			Status:       userCredential.Status,
			CreatedAt:    userCredential.CreatedAt,
			UpdatedAt:    userCredential.UpdatedAt,
			Deleted:      userCredential.Deleted,
		})

	}
	return response.Success(UserCredentialResponse, "Berhasil mengambil data")

}

// Function Search By Username
func (repo *UserRespository) SearchByUsername(username string) models.BaseResponse[[]models.SearchByUsernameResponse] {
	// Prepare Data
	var userCredential []entity.UserCredential
	var SearchByUsernameResponse []models.SearchByUsernameResponse
	var response = models.BaseResponse[[]models.SearchByUsernameResponse]{}

	err := repo.db.Where("username LIKE ?", "%"+username+"%").Find(&userCredential).Error
	if err != nil {
		return response.BadRequest(SearchByUsernameResponse, "User tidak ditemukan")
	}
	if len(userCredential) == 0 {
		return response.BadRequest(SearchByUsernameResponse, "User tidak ditemukan")
	}

	for _, userCredential := range userCredential {
		SearchByUsernameResponse = append(SearchByUsernameResponse, models.SearchByUsernameResponse{
			Id:       userCredential.ID,
			FullName: userCredential.FullName,
			UserName: userCredential.Username,
		})
	}

	return response.Success(SearchByUsernameResponse, "Berhasil mengambil data")
}
