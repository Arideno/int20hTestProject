package apiserver

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	router *gin.Engine
}

func NewServer() *APIServer {
	return &APIServer{router: gin.Default()}
}

func (s *APIServer) Start() error {
	s.configureRouter()
	return s.router.Run()
}

func (s *APIServer) configureRouter() {

	s.router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{"GET"},
		AllowHeaders: []string{"Origin","Content-Type"},
	}))

	v1 := s.router.Group("/v1")
	v1.GET("/info", handleInfo())
}

func handleInfo() gin.HandlerFunc {
	type response struct {
		Products []product
	}
	return func(c *gin.Context) {
		store := c.Query("store")

		switch store {
		case "atb":
			atb(func(products []product) {
				c.JSON(200, response{Products: products})
			})
			return
		case "silpo":
			c.JSON(200, response{Products: Silpo()})
			return
		case "auchan":
			c.JSON(200, response{Products: Auchan()})
			return

		}

		c.JSON(400, gin.H{
			"message": "no such store",
		})
	}
}
