package main

import (
	"github.com/Bluhabit/uwang-rest-account/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	//var err error
	//err = godotenv.Load()
	//if err != nil {
	//	log.Fatalf("Error getting env, %v", err)
	//} else {
	//	fmt.Println("we are getting the env values")
	//}
	//
	//contoh := os.Getenv("GIN_MODE")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Halo blue",
		})
	})

	routes.InitRoutes(r)

	if err := r.Run(":8000"); err != nil {
		log.Fatal("Gagal memulai server")
	}
}
