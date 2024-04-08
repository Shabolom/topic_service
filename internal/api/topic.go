package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
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

func (ta *TopicApi) Create(c *gin.Context) {
	var topic models.Topic

	_, pathToLogo, err := tools.MultipartFormTopic(c, &topic)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Error(err)
		return
	}

	err = topicService.Create(topic, pathToLogo)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Error(err)
		return
	}

	c.String(http.StatusCreated, "топик успешно создан")
	defer c.Request.Body.Close()
}

func (ta *TopicApi) Get(c *gin.Context) {
	topicID := c.Param("id")

	result, err := topicService.Get(topicID)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (ta *TopicApi) GetAll(c *gin.Context) {

	result, err := topicService.GetAll()
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (ta *TopicApi) Update(c *gin.Context) {
	var topic models.Topic
	topicID := c.Param("id")

	_, pathToLogo, err := tools.MultipartFormTopic(c, &topic)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Error(err)
		return
	}

	err = topicService.Update(topic, pathToLogo, topicID)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Error(err)
		return
	}

	c.String(http.StatusCreated, "топик упешно обновлен")
}
