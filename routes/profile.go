package routes

import (
	"github.com/Bluhabit/uwang-rest-account/common"
	"github.com/Bluhabit/uwang-rest-account/repositories"
	"github.com/gin-gonic/gin"
)

func UpdateProfileUsername(ctx *gin.Context) {
	// ambil request dari user

	claims, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(401, gin.H{
			"status_code": 401,
			"data":        nil,
			"message":     "Token not provided",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status_code": 401,
		"data":        claims,
		"message":     "Token not provided",
	})

	user := repositories.GetUserById(claims.(common.UserClaims).Sub)

	if user == nil {
		ctx.JSON(401, gin.H{
			"status_code": 401,
			"data":        claims,
			"message":     "Token not provided",
		})
	}

	//update here

}

func UpdateProfilePicture(ctx *gin.Context) {

}

func UpdateProfileInterestTopics(ctx *gin.Context) {

}

func UpdateProfileLevel(ctx *gin.Context) {

}
