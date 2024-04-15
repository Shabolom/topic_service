package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"service_topic/internal/models"
	"service_topic/internal/service"
	"service_topic/internal/tools"
)

type TopicApi struct {
}

func NewTopicApi() *TopicApi {
	return &TopicApi{}
}

var topicService = service.NewTopicService()

// Create создание топика
//
// @Summary	получение сообщений в конкретном топике
// @Security ApiKeyAuth
// @Accept	json
// @Produce	json
// @Tags	Topic
// @Param	data	query		models.Topic	true	"информация о топике"
// @Param	file	formData	file			false	"лого топика"
// @Success	200		{string}	string 	"топик успешно создан"
// @Failure	400		{object}	models.Error
// @Router	/api/topic [post]
func (ta *TopicApi) Create(c *gin.Context) {
	var topic models.Topic

	_, pathToLogo, err := tools.MultipartForm(c, &topic, "topic")
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	err = topicService.Create(topic, pathToLogo)
	if err != nil {
		if err2 := os.Remove(pathToLogo); err2 != nil {
			log.WithField("component", "api").Debug(err2)

		}
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.String(http.StatusCreated, "топик успешно создан")
	defer c.Request.Body.Close()
}

// Get получение информации о конкретном топике
//
// @Summary	получение информации о конкретном топике
// @Security ApiKeyAuth
// @Accept	json
// @Produce	json
// @Tags	Topic
// @Param	id		path		string	true	"id топика"
// @Success	200		{object}	domain.Topic
// @Failure	400		{object}	models.Error
// @Router	/api/topic/{id} [get]
func (ta *TopicApi) Get(c *gin.Context) {
	topicID := c.Param("id")

	result, err := topicService.Get(topicID)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Debug(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (ta *TopicApi) GetAll(c *gin.Context) {

	limit, skip, err := tools.Pagination(c)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	result, err := topicService.GetAll(limit, skip)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Update обновление топика
//
// @Summary	обновление топика
// @Security ApiKeyAuth
// @Accept	json
// @Produce	json
// @Tags	Topic
// @Param	id		path		string	true	"id топика"
// @Param	file	formData	file	false	"логотип"
// @Param	data	query		models.Topic	false	"информация о топике"
// @Success	200		{string}	string 	"топик упешно обновлен"
// @Failure	400		{object}	models.Error
// @Router	/api/topic/{id} [patch]
func (ta *TopicApi) Update(c *gin.Context) {
	var topic models.Topic
	topicID := c.Param("id")

	_, pathToLogo, err := tools.MultipartForm(c, &topic, "topic")
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Debug(err)
		return
	}

	err = topicService.Update(topic, pathToLogo, topicID)
	if err != nil {
		if err2 := os.Remove(pathToLogo); err2 != nil {
			log.WithField("component", "api").Debug(err2)
		}
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Debug(err)
		return
	}

	c.String(http.StatusCreated, "топик упешно обновлен")
}

// JoinTopic вступление в топик
//
// @Summary	вступление в топик
// @Security ApiKeyAuth
// @Accept	json
// @Produce	json
// @Tags	Topic
// @Param	id		path		string	true	"id топика"
// @Success	200		{string}	string	"вы присоединились к тоопику"
// @Failure	400		{object}	models.Error
// @Router	/api/topic/{id} [put]
func (ta *TopicApi) JoinTopic(c *gin.Context) {
	topicID := c.Param("id")

	claims, err := tools.ParseTokenClaims(c)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Debug(err)
		return
	}

	err = topicService.JoinTopic(topicID, claims.UserID)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.String(http.StatusCreated, "вы присоединились к тоопику")
}

// DeleteTopic удаление топика
//
// @Summary	получение сообщений в конкретном топике
// @Security ApiKeyAuth
// @Accept	json
// @Produce	json
// @Tags	Topic
// @Param	id		path		string	true	"id топика"
// @Success	200		{string}	string	"топик удален"
// @Failure	400		{object}	models.Error
// @Router	/api/topic/{id} [delete]
func (ta *TopicApi) DeleteTopic(c *gin.Context) {

	claims, err := tools.ParseTokenClaims(c)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "api").Debug(err)
		return
	}

	if claims.UserPerm != 3 {
		tools.CreateError(http.StatusBadRequest, errors.New("не достаточно прав"), c)
		return
	}

	topicID := c.Param("id")

	err = topicService.DeleteTopic(topicID)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.String(http.StatusCreated, "топик удален")
}

// @Summary	получение сообщений в конкретном топике
// @Security ApiKeyAuth
// @Accept	json
// @Produce	json
// @Tags	Topic
// @Param	topic_id	query		string	true	"id топика"
// @Param	id		query		string	true	"id пользователя"
// @Success	200		{string}	string	 "поьзователь удален из тоопика"
// @Failure	400		{object}	models.Error
// @Router	/api/topic/user [delete]
func (ta *TopicApi) DeleteUser(c *gin.Context) {
	topicID := c.Query("topic_id")
	userID := c.Query("id")

	err := tools.PermCheck(c)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	err = topicService.DeleteUser(topicID, userID)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.String(http.StatusCreated, "поьзователь удален из тоопика")
}

// TopicRating отображает рейтин топиков по количеству пользователей
//
// @Summary	отображает рейтин топиков по количеству пользователей
// @Security ApiKeyAuth
// @Accept	json
// @Produce	json
// @Tags	Topic
// @Param	limit	query		string	true	"количество элементов на странице"
// @Param	page	query		string	true	"страинца"
// @Success	200		{object}	[]models.TopicRating
// @Failure	400		{object}	models.Error
// @Router	/api/topic/rating [get]
func (ta *TopicApi) TopicRating(c *gin.Context) {

	limit, skip, err := tools.Pagination(c)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	result, err := topicService.TopicRating(limit, skip)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusCreated, result)
}
