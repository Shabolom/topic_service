package service

import (
	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"
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
	id, err := uuid.NewV4()
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return err
	}

	topicEntity := domain.Topic{
		ID:          id,
		WhenCreated: time.Now(),
		WhenUpdate:  time.Now(),
		TopicInfo:   topic.Info,
		TopicName:   topic.Name,
		TopicFile:   pathToFile,
	}

	err = topicRepo.Create(topicEntity)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TopicService) Get(topicID string) (domain.Topic, error) {
	id, err := uuid.FromString(topicID)
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return domain.Topic{}, err
	}

	result, err := topicRepo.FindOne("id", id)
	if err != nil {
		return domain.Topic{}, err
	}

	return result, nil
}

func (ts *TopicService) GetAll(limit, skip uint64) ([]domain.Topic, error) {
	result, err := topicRepo.GetAll(limit, skip)
	if err != nil {
		return []domain.Topic{}, err
	}

	return result, nil
}

func (ts *TopicService) Update(topic models.Topic, pathToFile, topicID string) error {
	id, err := uuid.FromString(topicID)
	if err != nil {
		log.WithField("component", "service").Debug(err)
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

func (ts *TopicService) JoinTopic(strTopicID string, userID uuid.UUID) error {
	topicID, err := uuid.FromString(strTopicID)
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return err
	}

	id, err := uuid.NewV4()
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return err
	}

	usersTopicsEntity := domain.UsersTopics{
		ID:          id,
		WhenCreated: time.Now(),
		UserID:      userID,
		TopicID:     topicID,
	}

	err = topicRepo.JoinTopic(usersTopicsEntity)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TopicService) DeleteTopic(strTopicID string) error {
	topicID, err := uuid.FromString(strTopicID)
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return err
	}

	err = topicRepo.DeleteTopic(topicID)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TopicService) DeleteUser(strTopicID, strUserID string) error {
	topicID, err := uuid.FromString(strTopicID)
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return err
	}

	userID, err := uuid.FromString(strUserID)
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return err
	}

	err = topicRepo.DeleteUser(topicID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TopicService) TopicRating(limit, skip uint64) ([]models.TopicRating, error) {

	result, err := topicRepo.TopicRating(limit, skip)
	if err != nil {
		return []models.TopicRating{}, err
	}

	return result, nil
}
