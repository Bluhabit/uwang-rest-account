package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.POST("/update-profile-username", UpdateProfileUsername)
		v1.POST("/update-profile-picture", UpdateProfilePicture)
		v1.POST("/update-profile-interest-topics", UpdateProfileInterestTopics)
		v1.POST("/update-profile-level", UpdateProfileLevel)
	}
}
