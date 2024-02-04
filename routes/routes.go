package routes

import (
	"github.com/Bluhabit/uwang-rest-account/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.POST("/update-profile-username", UpdateProfileUsername)
		v1.POST("/update-profile-picture", middlewares.AuthMiddleware(), UpdateProfilePicture)
		v1.POST("/update-profile-interest-topics", middlewares.AuthMiddleware(), UpdateProfileInterestTopics)
		v1.POST("/update-profile-level", middlewares.AuthMiddleware(), UpdateProfileLevel)
	}
}
