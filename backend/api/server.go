package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Server will serve HTTP requests
type Server struct {
	store  *sqlx.DB    // the database store
	router *gin.Engine // the routing for the http
}

// Creates a new HTTP server and setup routing
func NewServer(store *sqlx.DB) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/periods/:DateStamp", server.getPeriods)
	router.POST("/periods", server.postPeriods)
	router.PUT("/periods", server.putPeriods) // should be optimized
	router.DELETE("/periods/:ID", server.deletePeriods)
	router.GET("/achievements/:DateStamp", server.getAchievements)
	router.POST("/achievements", server.postAchievements)
	router.PUT("/achievements", server.putAchievements) // should be optimized
	router.DELETE("/achievements/:ID", server.deleteAchievements)

	server.router = router
	return server
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
