package v1

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nicolasassi/the-one-api/infrastructure"
)

const (
	path   = "/api/v1/"
	dbName = "the-one"
)

func Serve(port string) error {
	services := infrastructure.NewRepositories(infrastructure.WithMongoDB(nil, dbName))

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET"}
	config.AllowHeaders = []string{"Content-Type"}

	r.GET(path + "book/")
	return r.Run(":" + port)
}
