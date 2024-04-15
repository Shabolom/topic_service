package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"service_topic/internal/models"
	"service_topic/internal/service"
	"service_topic/internal/tools"
	"strings"
)

type MessagesApi struct {
}

func NewMessagesApi() *MessagesApi {
	return &MessagesApi{}
}

var messagesService = service.NewMessageService()

// Post оставить сообщение в определенном топике
//
// @Summary	оставить сообщение в определенном топике
// @Security ApiKeyAuth
// @Accept	json
// @Produce	json
// @Tags	Message
// @Param	id		path		string	true	"id топика"
// @Param	file	formData	file	false	"файлы"
// @Param	data	query		models.Messages	false	"сообщение"
// @Success	200		{object}	domain.Message
// @Failure	400		{object}	models.Error
// @Router	/api/message/{id} [get]
func (ma MessagesApi) Post(c *gin.Context) {
	var message models.Messages
	topicID := c.Param("id")

	claims, pathToFile, err := tools.MultipartForm(c, &message, "messages")
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	result, err := messagesService.Post(topicID, pathToFile, message.Message, claims.UserID)
	if err != nil {
		files := strings.Split(pathToFile, "(space)")
		for _, file := range files {
			if err2 := os.Remove(file); err2 != nil {
				log.WithField("component", "api").Debug(err)
			}
		}
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// Update редактирование/изменение сообщения
//
// @Summary	редактирование/изменение сообщения
// @Security ApiKeyAuth
// @Accept	json
// @Produce	json
// @Tags	Message
// @Param	id		path		string	true	"id топика"
// @Param	file	formData	file	false	"файлы"
// @Param	data	query		models.Messages	false	"сообщение"
// @Success	200		{object}	domain.Message
// @Failure	400		{object}	models.Error
// @Router	/api/message/{id} [patch]
func (ma MessagesApi) Update(c *gin.Context) {
	var message models.Messages
	messageID := c.Param("id")

	_, pathToFile, err := tools.MultipartForm(c, &message, "messages")
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	result, err := messagesService.Update(messageID, pathToFile, message.Message)
	if err != nil {
		files := strings.Split(pathToFile, "(space)")
		for _, file := range files {
			if err2 := os.Remove(file); err2 != nil {
				log.WithField("component", "api").Debug(err)
			}
		}
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// Get получение все сообщений в топике
//
// @Summary	получение все сообщений в топике
// @Security ApiKeyAuth
// @Accept	json
// @Produce	json
// @Tags	Message
// @Param	id		path		string	true	"id топика"
// @Success	200		{object}	[]models.RespMessage
// @Failure	400		{object}	models.Error
// @Router	/api/message/{id} [get]
func (ma MessagesApi) Get(c *gin.Context) {
	topicID := c.Param("id")

	result, err := messagesService.Get(topicID)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Delete удаление сообщения
//
// @Summary	удаление сообщения
// @Security ApiKeyAuth
// @Accept	json
// @Produce	json
// @Tags	Message
// @Param	id		path		string	true	"id сообщения"
// @Success	200		{string}	string ""
// @Failure	400		{object}	models.Error
// @Router	/api/message/{id} [delete]
func (ma MessagesApi) Delete(c *gin.Context) {
	messageID := c.Param("id")

	calims, err := tools.ParseTokenClaims(c)
	if err != nil {
		log.WithField("component", "service").Debug(err)
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	if calims.UserPerm != 3 {
		err = messagesService.Delete(calims.UserID, messageID)
		if err != nil {
			tools.CreateError(http.StatusBadRequest, err, c)
			return
		}

		c.Status(http.StatusNoContent)
		return
	}

	err = messagesService.Delete(uuid.Nil, messageID)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.Status(http.StatusNoContent)
}

// RatingMessages показывает рейтинг сообщений по лайкам в конкретном топике
//
// @Summary	показывает рейтинг сообщений по лайкам в конкретном топике
// @Security ApiKeyAuth
// @Accept	json
// @Produce	json
// @Tags	Message
// @Param	id		path		string	true	"id топика"
// @Param	limit	query		string	true	"количество элементов"
// @Param	page	query		string	true	"страница"
// @Success	200		{object}	[]models.RespMessage
// @Failure	400		{object}	models.Error
// @Router	/api/message/topic/{id} [get]
func (ma MessagesApi) RatingMessages(c *gin.Context) {
	topicID := c.Param("id")

	limit, skip, err := tools.Pagination(c)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	result, err := messagesService.RatingMessages(topicID, limit, skip)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusNoContent, result)
}

// Rating оставляем оценку к сообщению
//
// @Summary	оставляем оценку к сообщению
// @Security ApiKeyAuth
// @Accept	json
// @Produce	json
// @Tags	Message
// @Param	id		path		string	true	"id сообщения"
// @Param	like	query		string	true	"true/false лайк или диз лайк"
// @Success	200		{string}	string 	"оценка изменена"
// @Failure	400		{object}	models.Error
// @Router	/api/message/rating/{id} [get]
func (ma MessagesApi) Rating(c *gin.Context) {
	messageID := c.Param("id")
	like := c.Query("like")

	claims, err := tools.ParseTokenClaims(c)
	if err != nil {
		log.WithField("component", "api").Debug(err)
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	err = messagesService.Rating(like, messageID, claims.UserID)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, "оценка изменена")
}
