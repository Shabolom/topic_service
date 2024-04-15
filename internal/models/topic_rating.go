package models

import "service_topic/internal/domain"

type TopicRating struct {
	Topic domain.Topic
	Users int
}
