package middlware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"service_topic/internal/tools"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := tools.ParseTokenClaims(c)
		if err != nil {
			tools.CreateError(http.StatusBadRequest, err, c)
			log.WithField("component", "middleware").Debug(err)
			return
		}

		if claims.UserStatus != "confirmed" {
			tools.CreateError(http.StatusBadRequest, err, c)
			log.WithField("component", "middleware").Debug(err)
			return
		}

		c.Next()
	}
}
