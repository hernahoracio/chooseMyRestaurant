package server

import (
	"net/http"

	"chooseMyRestaurant/repo"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Repo repo.Repo
}

func (s *Server) Engine() *gin.Engine {
	eng := gin.New()
	eng.Use(gin.Recovery())

	//montoring stuff
	eng.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })

	//external API for GPS coordinates and restaurant locations
	api := eng.Group("/")
	restaurantsHandler := api.Group("/restaurants")
	restaurantsHandler.GET("")

	return eng
}
