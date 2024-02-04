package routes

import (
	"github.com/Bluhabit/uwang-rest-account/common"
	"github.com/Bluhabit/uwang-rest-account/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateProfileUsername(ctx *gin.Context) {
	// ambil request dari user

	//claims, exists := ctx.Get("user")
	//if !exists {
	//	ctx.JSON(401, gin.H{
	//		"status_code": 401,
	//		"data":        nil,
	//		"message":     "Token not provided",
	//	})
	//	return
	//}
	//
	//ctx.JSON(200, gin.H{
	//	"status_code": 401,
	//	"data":        claims,
	//	"message":     "Token not provided",
	//})
	//
	//user := repositories.GetUserById(claims.(common.UserClaims).Sub)
	//
	//if user == nil {
	//	ctx.JSON(401, gin.H{
	//		"status_code": 401,
	//		"data":        claims,
	//		"message":     "Token not provided",
	//	})
	//}

	// Ambil request dari user
	var updateRequest struct {
		Username string `json:"username"`
	}

	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	//claims, exists := ctx.Get("user")
	claims, exists := common.DummyToken("0ea085da-b618-42a1-8130-019195bf5e81")
	if !exists {
		ctx.JSON(401, gin.H{
			"status_code": 401,
			"data":        nil,
			"message":     "Token not Provided",
		})
		return
	}

	userCredential := repositories.GetUserProfileById("0ea085da-b618-42a1-8130-019195bf5e81")

	if userCredential == nil {
		ctx.JSON(401, gin.H{
			"status_code": 401,
			"data":        claims,
			"message":     "Profil user gak ketemu",
		})
		return
	}

	// Update username
	userCredential.Username = updateRequest.Username
	repositories.UpdateUserProfile(userCredential)

	ctx.JSON(200, gin.H{
		"status_code": 200,
		"data":        userCredential,
		"message":     "Username updated successfully",
	})
}

func UpdateProfilePicture(ctx *gin.Context) {

}

func UpdateProfileInterestTopics(ctx *gin.Context) {

}

func UpdateProfileLevel(ctx *gin.Context) {

}
