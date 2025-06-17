package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zeusro/system/internal/core/logprovider"
	"github.com/zeusro/system/internal/core/webprovider"
	"github.com/zeusro/system/internal/service"
)

type IndexRoutes struct {
	logger logprovider.Logger
	gin    webprovider.MyGinEngine
	s      service.HealthService
	// m middleware.JWTMiddleware
}

func NewIndexRoutes(logger logprovider.Logger, gin webprovider.MyGinEngine, s service.HealthService) IndexRoutes {
	return IndexRoutes{
		logger: logger,
		gin:    gin,
		s:      s,
	}
}

func (r IndexRoutes) SetUp() {

	r.gin.Gin.GET("/index", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	index := r.gin.Gin.Group("/api")
	{
		//http://localhost:8080/api/health
		index.OPTIONS("health", r.s.Check)
		index.GET("health", r.s.Check)
		index.OPTIONS("healthz", r.s.Check)
		index.GET("healthz", r.s.Check)
	}

}
