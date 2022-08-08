package api

import (
	"github.com/arioki1/qiscus-caa/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type API struct {
	Conf config.Config
}

func NewAPI(conf config.Config) API {
	return API{
		Conf: conf,
	}
}

func (a *API) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func (a *API) RouteNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "resource not found",
	})
}

func (a *API) AbortWebhookHandling(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"ok":      false,
		"message": message,
	})
}
