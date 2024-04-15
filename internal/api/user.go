package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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

// Register регистрация пользователия через сторонний сервис
//
// @Summary	регистрация пользователия через сторонний сервис
// @Accept	json
// @Produce	json
// @Tags	User
// @Param	ввод	body		models.User		true	"логин и пароль"
// @Success	201		{string}	string 	"успешно зарегестрировались"
// @Failure	400		{object}	models.Error
// @Router	/api/user/register [post]
func (ua *UserAPI) Register(c *gin.Context) {
	var user models.User

	response, err := tools.CreateRequest(c, "POST", config.Env.ConnectionApi)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Debug(err)
		return
	}

	if response.StatusCode != 201 {
		var respErr models.RespError

		err2 := tools.ShortUnmarshal(response.Body, &respErr)
		if err2 != nil {
			tools.CreateError(http.StatusBadRequest, err2, c)
			log.WithField("component", "rest").Debug(err2)
			return
		}

		c.JSON(response.StatusCode, respErr)
		return
	}

	err = tools.ShortUnmarshal(response.Body, &user)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Debug(err)
		return
	}

	err = userService.Register(user)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Debug(err)
		return
	}

	c.JSON(http.StatusCreated, "успешно зарегестрировались")
}

// Login авторизация пользователия через сторонний сервис
//
// @Summary	авторизация пользователия через сторонний сервис
// @Accept	json
// @Produce	json
// @Tags	User
// @Param	ввод	body		models.User		true	"логин и пароль"
// @Success	201		{string}	string 	"вы успешно авторезировались"
// @Failure	400		{object}	models.Error
// @Router	/api/user/login [post]
func (ua *UserAPI) Login(c *gin.Context) {
	response, err := tools.CreateRequest(c, "POST", config.Env.ConnectionLogin)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Debug(err)
		return
	}

	if response.StatusCode != 200 {
		var respErr models.RespError

		err2 := tools.ShortUnmarshal(response.Body, &respErr)
		if err2 != nil {
			tools.CreateError(http.StatusBadRequest, err2, c)
			log.WithField("component", "rest").Debug(err2)
			return
		}

		c.JSON(response.StatusCode, respErr)
		return
	}

	c.JSON(http.StatusOK, "вы успешно авторезировались")
}
