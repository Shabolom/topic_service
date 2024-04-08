package middlware

import (
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"service_topic/internal/tools"
)

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := tools.ParseTokenClaims(c)
		if err != nil {
			tools.CreateError(http.StatusBadRequest, err, c)
			log.WithField("component", "rest").Error(err)
			return
		}

		if claims.UserStatus != "confirmed" {
			tools.CreateError(http.StatusBadRequest, err, c)
			log.WithField("component", "rest").Error(err)
			return
		}

		c.Next()
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := tools.ParseTokenClaims(c)
		if err != nil {
			tools.CreateError(http.StatusBadRequest, err, c)
			log.WithField("component", "rest").Error(err)
			return
		}

		if claims.UserPerm != 3 {
			tools.CreateError(http.StatusBadRequest, errors.New("не хватает прав доступа"), c)
			log.WithField("component", "rest").Error(errors.New("не хватает прав доступа"))
			return
		}

		c.Next()
	}
}
