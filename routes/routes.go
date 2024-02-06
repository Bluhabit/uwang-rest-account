package routes

import (
	"github.com/Bluhabit/uwang-rest-account/middlewares"
	"github.com/Bluhabit/uwang-rest-account/routes/profile"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.POST("/update-profile-username", middlewares.AuthMiddleware(), profile.UpdateProfileUsername)
		v1.POST("/update-profile-picture", middlewares.AuthMiddleware(), profile.UpdateProfilePicture)
		v1.POST("/update-profile-interest-topics", middlewares.AuthMiddleware(), profile.UpdateProfileInterestTopics)
		v1.POST("/update-profile-level", middlewares.AuthMiddleware(), profile.UpdateProfileLevel)
	}
}
