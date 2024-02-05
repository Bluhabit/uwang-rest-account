package routes

import (
	"github.com/Bluhabit/uwang-rest-account/routes/profile"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.POST("/update-profile-username", profile.UpdateProfileInterestTopics)
		v1.POST("/update-profile-picture", profile.UpdateProfileInterestTopics)
		v1.POST("/update-profile-interest-topics", profile.UpdateProfileInterestTopics)
		v1.POST("/update-profile-level", profile.UpdateProfileInterestTopics)
	}
}
