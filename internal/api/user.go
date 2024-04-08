package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"service_topic/config"
	"service_topic/internal/models"
	"service_topic/internal/service"
	"service_topic/internal/tools"
)

type UserAPI struct {
}

func NewUserApi() *UserAPI {
	return &UserAPI{}
}

var userService = service.NewUserService()

func (ua *UserAPI) Register(c *gin.Context) {
	var user models.User

	response, err := tools.CreateRequest(c, "POST", config.Env.ConnectionApi)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Error(err)
		return
	}

	if response.StatusCode != 201 {
		data, err2 := io.ReadAll(response.Body)
		str := tools.StringFormat(string(data))
		if err2 != nil {
			tools.CreateError(http.StatusBadRequest, err2, c)
			log.WithField("component", "rest").Error(err2)
			return
		}

		c.JSON(response.StatusCode, str)
		return
	}

	err = tools.ShortUnmarshal(response.Body, &user)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Error(err)
		return
	}

	err = userService.Register(user)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Error(err)
		return
	}

	c.JSON(http.StatusCreated, "успешно зарегестрировались")
}

func (ua *UserAPI) Login(c *gin.Context) {
	response, err := tools.CreateRequest(c, "POST", config.Env.ConnectionLogin)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Error(err)
		return
	}

	if response.StatusCode != 200 {
		data, err2 := io.ReadAll(response.Body)
		str := tools.StringFormat(string(data))
		if err2 != nil {
			tools.CreateError(http.StatusBadRequest, err2, c)
			log.WithField("component", "rest").Error(err2)
			return
		}

		c.JSON(response.StatusCode, str)
		return
	}

	c.JSON(http.StatusOK, "вы успешно авторезировались")
}
