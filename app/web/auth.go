package web

import (
	"github.com/gin-gonic/gin"
	"github.com/pieterclaerhout/go-log"
	"github.com/voltento/mongo_rest/app/config"
	"net/http"
)

func basicAuth(cfg config.Config, c *gin.Context) {
	user, password, hasAuth := c.Request.BasicAuth()
	if hasAuth && user == cfg.User && password == cfg.Password {
		log.Debugf("user %v authenticated", user)
	} else {
		c.JSON(http.StatusUnauthorized, map[string]string{"status": "not authorized"})
		c.Abort()
		return
	}
}
