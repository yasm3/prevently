package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/yasm3/prevently/internal/db"
	"github.com/yasm3/prevently/internal/logger"
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

	a.Router.GET("/users", func(c *gin.Context) {
		u, _ := uuid.NewV4()
		var pgid pgtype.UUID
		_ = pgid.Scan(u.String())
		res, err := a.DB.GetUserByID(c, pgid)
		if err != nil {
			c.JSON(404, err.Error())
		}
		c.JSON(200, res)
	})
}

func (a *APIServer) Run() {
	a.Logger.Info("Starting server")
	a.Router.Run(":" + a.Port)
}
