package main

import (
	"fmt"
	"gin_template/src/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	config.InitEnv()
	// database.ConnectDB()

	api := r.Group("/api")
	v1 := api.Group("/v1")
	fmt.Print(v1)

	PORT := config.GetEnv("PORT", "8000")

	r.Run(fmt.Sprintf(":%s", PORT))
}
