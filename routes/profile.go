package routes

import (
	"github.com/Bluhabit/uwang-rest-account/common"
	"github.com/Bluhabit/uwang-rest-account/entity"
	"github.com/Bluhabit/uwang-rest-account/models"
	"github.com/Bluhabit/uwang-rest-account/repositories"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

func UpdateProfileUsername(ctx *gin.Context) {
	// Ambil request dari user
	var response = models.BaseResponse[string]{}

	var updateRequest models.UpdateProfileUsername

	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(200, response.BadRequest("", err.Error()))
	}

	// cek username yang ada
	existingUser := repositories.GetUserByUsername(updateRequest.Username)
	if existingUser != nil {
		ctx.JSON(200, response.BadRequest("", "username sudah digunakan"))
	}

	//claims, exists := ctx.Get("user")
	claims, exists := common.DummyToken("0ea085da-b618-42a1-8130-019195bf5e81")
	if !exists {
		ctx.JSON(200, response.BadRequest("", "Token not Provided"))
		return
	}

	userCredential := repositories.GetUserProfileById(claims)

	if userCredential == nil {
		ctx.JSON(200, response.BadRequest("", "profil user tidak ditemukan"))
		return
	}

	// Update username
	userCredential.Username = updateRequest.Username
	repositories.UpdateUserProfile(userCredential)

	ctx.JSON(200, response.Success("", "Berhasil merubah username"))
}

func UpdateProfilePicture(ctx *gin.Context) {
	var response = models.BaseResponse[string]{}
	var updateRequest models.UpdateProfilePicture

	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(200, response.BadRequest("", err.Error()))
	}

	//userid, exists := ctx.Get("user")
	userid, exists := common.DummyToken("0ea085da-b618-42a1-8130-019195bf5e81")
	if !exists {
		ctx.JSON(200, response.BadRequest("", "Token not Provided"))
		return
	}

	userProfile := repositories.GetProfilePictureByUserID(userid)

	var profilePictureID = uuid.NewString()

	if userProfile == nil {
		// create data
		newProfile := &entity.UserProfile{
			ID:        profilePictureID,
			Key:       "profile-picture",
			Value:     updateRequest.ProfilePicture,
			UserID:    userid,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Deleted:   false,
		}

		err := repositories.CreateUserProfilePicture(newProfile)
		if err != nil {
			ctx.JSON(200, response.BadRequest("", "foto profil gagal disimpan"))
			return
		}
		ctx.JSON(200, response.Success("", "berhasil membuat foto profil"))
		return
	}

	userProfile.Value = updateRequest.ProfilePicture
	repositories.UpdateUserProfilePicture(userProfile)
	ctx.JSON(200, response.Success("", "berhasil menyimpan foto profil"))
}

func UpdateProfileInterestTopics(ctx *gin.Context) {

}

func UpdateProfileLevel(ctx *gin.Context) {

}
