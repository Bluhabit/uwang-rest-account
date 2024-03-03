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
		v1.POST("/update-profile")

		v1.GET("/get-detail-user", middlewares.AuthMiddleware(), profile.DetailUserResponse)
		v1.GET("/profile", middlewares.AuthMiddleware(), profile.GetDetailProfile)
	}

	v1Admin := router.Group("/v1/admin")
	{
		v1Admin.GET("/get-profile", middlewares.AuthMiddleware(), profile.GetDetailProfile)
		v1Admin.GET("/get-statistic")
		v1Admin.GET("/get-top-user")
		v1Admin.GET("/")
	}
}
