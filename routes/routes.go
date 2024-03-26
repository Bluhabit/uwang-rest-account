package routes

import (
	"github.com/Bluhabit/uwang-rest-account/middlewares"
	"github.com/Bluhabit/uwang-rest-account/routes/profile"
	"github.com/Bluhabit/uwang-rest-account/routes/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/profile", middlewares.AuthMiddleware(), profile.GetProfile)
		v1.POST("/update-profile-username", middlewares.AuthMiddleware(), profile.UpdateProfileUsername)
		v1.POST("/update-profile-picture", middlewares.AuthMiddleware(), profile.UpdateProfilePicture)
		v1.POST("/update-profile-interest-topics", middlewares.AuthMiddleware(), profile.UpdateProfileInterestTopics)
		v1.POST("/update-profile-level", middlewares.AuthMiddleware(), profile.UpdateProfileLevel)
		v1.POST("/update-profile", middlewares.AuthMiddleware(), profile.UpdateProfile)
	}

	v1Admin := router.Group("/v1/admin")
	{
		v1Admin.GET("/search-by-username", middlewares.AuthMiddleware(), user.SearchByUsername)
		v1Admin.GET("/get-list-user", middlewares.AuthMiddleware(), user.GetListUserWithPaginate)
		v1Admin.GET("/get-detail-user", middlewares.AuthMiddleware(), user.GetDetailUser)
		v1Admin.GET("/get-profile", middlewares.AuthMiddleware(), profile.GetProfile)
		v1Admin.GET("/get-statistic", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"status_code": 200,
				"data": gin.H{
					"active_user": 0,
					"user_claim":  0,
					"total_user":  0,
				},
				"message": "Data statistik",
			})
		})
		v1Admin.GET("/get-top-user", middlewares.AuthMiddleware(), user.GetTopSevenUser)
	}