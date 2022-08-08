package router

import (
	"github.com/arioki1/qiscus-caa/api"
	"github.com/arioki1/qiscus-caa/config"
	"github.com/arioki1/qiscus-caa/src/api/v1/middleware"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func NewRouter(conf config.Config) *gin.Engine {
	r := gin.New()

	r.Use(requestid.New())
	r.Use(middleware.HTTPReqLog())
	r.Use(middleware.CORSMiddleware())

	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			log.Error().Msg(err)
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong on our side",
		})

		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	httpAPI := api.NewAPI(conf)

	r.GET("/", httpAPI.Home)

	r.NoRoute(httpAPI.RouteNotFound)

	return r
}
