package user

import (
	"github.com/Bluhabit/uwang-rest-account/models"
	"github.com/Bluhabit/uwang-rest-account/repositories/user"
	"github.com/gin-gonic/gin"
)

func GetDetailUser(ctx *gin.Context) {
	var response = models.BaseResponse[string]{}

	id, ok := ctx.GetQuery("detail")

	if !ok {
		ctx.JSON(200, response.BadRequest("", "User tidak ditemukan"))
	}

	repo := user.Init()
	data := repo.GetDetailUser(id)
	ctx.JSON(200, data)
}

func GetTopSevenUser(ctx *gin.Context) {
	repo := user.Init()
	data := repo.GetTopUser()
	ctx.JSON(200, data)
}

func GetListUserWithPaginate(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	size := ctx.DefaultQuery("size", "1")

	repo := user.Init()
	data := repo.GetListUser(page, size)
	ctx.JSON(200, data)
}

func SearchByUsername(ctx *gin.Context) {
	var response = models.BaseResponse[string]{}

	username, ok := ctx.GetQuery("username")

	if !ok {
		ctx.JSON(200, response.BadRequest("", "User tidak ditemukan"))
		return
	}

	repo := user.Init()
	data := repo.SearchByUsername(username)
	ctx.JSON(200, data)
}
