package apiserver

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"sort"
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
	v1.GET("/top", handleTop())
}

func handleTop() gin.HandlerFunc {
	type response struct {
		Products []product
	}
	return func(c *gin.Context) {
		atb(func(atbProducts []product) {
			allProducts := make([]product, 0)
			for _, p := range atbProducts {
				allProducts = append(allProducts, p)
			}

			silpoProducts := Silpo()
			for _, p := range silpoProducts {
				allProducts = append(allProducts, p)
			}

			auchanProducts := Auchan()
			for _, p := range auchanProducts {
				allProducts = append(allProducts, p)
			}

			sort.Slice(allProducts, func(i, j int) bool {
				return allProducts[i].Price < allProducts[j].Price
			})

			c.JSON(200, allProducts[:3])
		})
	}
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
