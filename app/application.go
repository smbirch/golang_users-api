package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {
	//sets up api routes
	mapUrls()
	router.Run(":8080")
}
