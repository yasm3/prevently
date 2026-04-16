package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yasm3/prevently/internal/db"
	"github.com/yasm3/prevently/internal/http/handler"
	"github.com/yasm3/prevently/internal/logger"
	"github.com/yasm3/prevently/internal/service"
)

type APIServer struct {
	Port   string
	Router *gin.Engine
	DB     *db.Queries
	Logger *logger.Logger
}

func NewServer(db *db.Queries, l *logger.Logger) *APIServer {
	r := gin.New()

	server := APIServer{
		Port:   "8000",
		Router: r,
		DB:     db,
		Logger: l,
	}

	r.Use(gin.Recovery())
	r.Use(l.GinMiddleware())

	server.registerRoutes()

	return &server
}

func (a *APIServer) registerRoutes() {
	a.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userService := service.NewUserService(a.DB)
	userHandler := handler.NewUserHandler(userService)

	a.Router.GET("/users", userHandler.GetUser)
	a.Router.POST("/users", userHandler.CreateUser)
}

func (a *APIServer) Run() {
	a.Logger.Info("Starting server")
	a.Router.Run(":" + a.Port)
}
