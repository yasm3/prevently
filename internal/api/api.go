package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yasm3/prevently/internal/db"
	"github.com/yasm3/prevently/internal/logger"
)

type APIServer struct {
	Port   string
	Router *gin.Engine
	DB     *db.Queries
	Logger *logger.Logger
}

func NewServer(r *gin.Engine, q *db.Queries, l *logger.Logger) *APIServer {
	server := APIServer{
		Port:   "8000",
		Router: r,
		DB:     q,
		Logger: l,
	}

	return &server
}

func NewRouter(l *logger.Logger) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(l.GinMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}

func (a *APIServer) Run() {
	a.Logger.Info("Starting server")
	a.Router.Run(":" + a.Port)
}
