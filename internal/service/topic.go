package service

import (
	"github.com/gofrs/uuid"
	"service_topic/internal/domain"
	"service_topic/internal/models"
	"service_topic/internal/repository"
	"time"
)

type TopicService struct {
}

func NewTopicService() *TopicService {
	return &TopicService{}
}

var topicRepo = repository.NewTopicRepo()

func (ts *TopicService) Create(topic models.Topic, pathToFile string) error {
	id, _ := uuid.NewV4()

	topicEntity := domain.Topic{
		ID:          id,
		WhenCreated: time.Now(),
		TopicInfo:   topic.Info,
		TopicName:   topic.Name,
		TopicFile:   pathToFile,
	}

	err := topicRepo.Create(topicEntity)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TopicService) Get(topicID string) (domain.Topic, error) {
	id, err := uuid.FromString(topicID)
	if err != nil {
		return domain.Topic{}, err
	}

	result, err := topicRepo.FindOne("id", id)
	if err != nil {
		return domain.Topic{}, err
	}

	return result, nil
}

func (ts *TopicService) GetAll() ([]domain.Topic, error) {
	result, err := topicRepo.GetAll()
	if err != nil {
		return []domain.Topic{}, err
	}

	return result, nil
}

func (ts *TopicService) Update(topic models.Topic, pathToFile, topicID string) error {
	id, err := uuid.FromString(topicID)
	if err != nil {
		return err
	}

	topicEntity := domain.Topic{
		WhenUpdate: time.Now(),
		TopicInfo:  topic.Info,
		TopicName:  topic.Name,
		TopicFile:  pathToFile,
	}

	err = topicRepo.Update(topicEntity, id)
	if err != nil {
		return err
	}

	return nil
}
