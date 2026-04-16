package api

import "github.com/gin-gonic/gin"

type APIServer struct {
	Port   string
	Router *gin.Engine
}

func NewAPIServer(router *gin.Engine) *APIServer {
	server := APIServer{
		Port:   "8080",
		Router: router,
	}
	return &server
}

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}

func (a *APIServer) Run() {
	a.Router.Run(":" + a.Port)
}
