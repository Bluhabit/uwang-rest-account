package profile

import (
	"fmt"

	"github.com/Bluhabit/uwang-rest-account/models"
	"github.com/Bluhabit/uwang-rest-account/repositories/profile"
	"github.com/gin-gonic/gin"
)

func UpdateProfileUsername(ctx *gin.Context) {
	// Ambil request dari user
	var response = models.BaseResponse[string]{}
	var updateRequest models.UpdateProfileUsername

	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(200, response.BadRequest("", err.Error()))
	}
	//ambil id dari token
	sessionId := ctx.GetString("session_id")
	fmt.Println(sessionId)
	if len(sessionId) < 1 {
		ctx.JSON(401, response.BadRequest("", "Token not Provided"))
		return
	}
	repositories := profile.Init()
	processUpdate := repositories.UpdateProfileUsername(sessionId, updateRequest.Username)
	ctx.JSON(200, processUpdate)
}

func UpdateProfilePicture(ctx *gin.Context) {
	var response = models.BaseResponse[string]{}
	var updateRequest models.UpdateProfilePicture

	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(200, response.BadRequest("", err.Error()))
		return
	}
	//ambil id dari token
	sessionId := ctx.GetString("session_id")
	if len(sessionId) < 1 {
		ctx.JSON(200, response.BadRequest("", "Token not Provided"))
		return
	}
	repositories := profile.Init()
	processUpdate := repositories.UpdateProfilePicture(sessionId, updateRequest.ProfilePicture)
	ctx.JSON(200, processUpdate)
}

func UpdateProfileInterestTopics(ctx *gin.Context) {
	var response = models.BaseResponse[string]{}
	var updateRequest models.UpdateProfileInterestTopic

	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(200, response.BadRequest("", err.Error()))
		return
	}
	//ambil id dari token
	sessionId := ctx.GetString("session_id")
	if len(sessionId) < 1 {
		ctx.JSON(200, response.BadRequest("", "Token not Provided"))
		return
	}
	repositories := profile.Init()
	processUpdate := repositories.UpdateProfileTopics(sessionId, updateRequest.InterestTopic)
	ctx.JSON(200, processUpdate)
}

func UpdateProfileLevel(ctx *gin.Context) {
	var response = models.BaseResponse[string]{}
	var updateRequest models.UpdateProfileLevel

	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(200, response.BadRequest("", err.Error()))
		return
	}
	//ambil id dari token
	sessionId := ctx.GetString("session_id")
	if len(sessionId) < 1 {
		ctx.JSON(200, response.BadRequest("", "Token not Provided"))
		return
	}
	repositories := profile.Init()
	processUpdate := repositories.UpdateProfileLevel(sessionId, updateRequest.Level)
	ctx.JSON(200, processUpdate)
}

func DetailUserResponse(ctx *gin.Context) {
	var response = models.BaseResponse[string]{}

	id, ok := ctx.GetQuery("detail")

	if !ok {
		ctx.JSON(200, response.BadRequest("", "User tidak ditemukan"))
	}

	repo := profile.Init()
	data := repo.GetAllDetailUser(id)
	ctx.JSON(200, data)
}
