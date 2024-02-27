package user

import (
	"github.com/Bluhabit/uwang-rest-account/models"
	"github.com/Bluhabit/uwang-rest-account/repositories/profile"
	"github.com/gin-gonic/gin"
)

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