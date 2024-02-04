package routes

import (
	"github.com/Bluhabit/uwang-rest-account/common"
	"github.com/Bluhabit/uwang-rest-account/models"
	"github.com/Bluhabit/uwang-rest-account/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateProfileUsername(ctx *gin.Context) {
	// Ambil request dari user
	var response = models.BaseResponse[string]{}

	var updateRequest struct {
		Username string `json:"username"`
	}

	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
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

}

func UpdateProfileInterestTopics(ctx *gin.Context) {

}

func UpdateProfileLevel(ctx *gin.Context) {

}
