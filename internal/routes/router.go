package routes

import (
	"github.com/gin-gonic/gin"
	"service_topic/internal/api"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	user := api.NewUserApi()
	topic := api.NewTopicApi()

	r.POST("/api/user/register", user.Register)
	r.POST("/api/user/login", user.Login)
	r.GET("/api/topic", topic.GetAll)
	return r
}
