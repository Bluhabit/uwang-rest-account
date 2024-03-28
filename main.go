package main

import (
	"github.com/Bluhabit/uwang-rest-account/common"
	"github.com/Bluhabit/uwang-rest-account/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"time"
)

func main() {

	common.GenerateEntity()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		expirationTime := time.Now().Add(5 * time.Hour)
		token := common.EncodeJWT(common.UserClaims{
			Id:  "Trian",
			Sub: "Trian",
			Iat: time.Now().UnixMilli(),
			Exp: time.Now().Add(time.Duration(time.Duration.Hours(1))).UnixMilli(),
			RegisteredClaims: &jwt.RegisteredClaims{
				ID:        "Trian",
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		})

		decode := common.DecodeJWT("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJibHVoYWJpdC5pZCIsInN1YiI6ImFkMWUzNGRhLTA3MzYtNDk1Ni1hNTk4LTUyODU5OTI2ZWQzNiIsImlhdCI6MTcwNzEyNTA4MiwiZXhwIjoxNzA3MTI4NjgyfQ.GFZ0dXrd4HISvTMLJJh8POfsAtdpSjJmDBObkCfW4dQ")
		// decode := common.DecodeJWT("yJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJibHVoYWJpdC5pZCIsInN1YiI6InRyaWFuZGFtYWkiLCJpYXQiOjE3MDY5NDUwNzksImV4cCI6MTcwNjk0NTA3OX0.5oYF3TlI_VlKwR-YXNLH7O7PcSfsetR8yYDtmZOwXPY")

		c.JSON(http.StatusOK, gin.H{
			"Message": "Halo blue",
			"Token":   token,
			"Decode":  decode,
		})
	})

	routes.InitRoutes(r)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://admin-uwang.bluhabit.id", "http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://admin-uwang/bluhabit.id" || origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	if err := r.Run(":8000"); err != nil {
		log.Fatal("Gagal memulai server")
	}
}
