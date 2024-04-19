package routes

import (
	"github.com/gin-gonic/gin"
	"service_topic/internal/api"
	"service_topic/internal/middlware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	Auth := r.Group("/", middlware.Auth())

	user := api.NewUserApi()
	topic := api.NewTopicApi()
	message := api.NewMessagesApi()

	r.POST("/api/user/register", user.Register)
	r.POST("/api/user/login", user.Login)
	r.GET("/api/message/:id", message.Get)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	{
		Auth.POST("/api/topic", topic.Create)
		Auth.PATCH("/api/topic/:id", topic.Update)
		Auth.DELETE("/api/topic/:id", topic.DeleteTopic)
		Auth.PUT("/api/topic/:id", topic.JoinTopic)
		Auth.GET("/api/topic", topic.GetAll)
		Auth.GET("/api/topic/:id", topic.Get)
		Auth.DELETE("/api/topic/user", topic.DeleteUser)
		Auth.GET("/api/topic/rating", topic.TopicRating)

		Auth.POST("/api/message/:id", message.Post)
		Auth.PATCH("/api/message/:id", message.Update)
		Auth.DELETE("/api/message/:id", message.Delete)
		Auth.GET("/api/message/rating/:id", message.Rating)
		Auth.GET("/api/message/topic/:id", message.RatingMessages)
	}

	return r
}
