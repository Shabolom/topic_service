package service

import (
	"errors"
	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"
	"service_topic/internal/domain"
	"service_topic/internal/models"
	"service_topic/internal/repository"
	"strconv"
	"time"
)

type MessageService struct {
}

func NewMessageService() *MessageService {
	return &MessageService{}
}

var messageRepo = repository.NewMessageRepo()

func (ms *MessageService) Post(strTopicID, filePath, message string, userID uuid.UUID) (domain.Message, error) {

	messageID, err := uuid.NewV4()
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return domain.Message{}, err
	}

	topicID, err := uuid.FromString(strTopicID)
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return domain.Message{}, err
	}

	if err = messageRepo.CheckSubscribe(userID, topicID); err != nil {
		return domain.Message{}, errors.New("вы не являетесь участником топика")
	}

	messageEntity := domain.Message{
		ID:           messageID,
		WhenCreated:  time.Now(),
		WhenUpdated:  time.Now(),
		UserMessage:  message,
		UserFilePath: filePath,
		UserID:       userID,
		TopicID:      topicID,
	}

	result, err := messageRepo.Post(messageEntity)
	if err != nil {
		return domain.Message{}, err
	}

	return result, nil
}

func (ms *MessageService) Update(strMessageID, filePath, message string) (domain.Message, error) {

	messageID, err := uuid.FromString(strMessageID)
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return domain.Message{}, err
	}


	messageEntity := domain.Message{
		ID:           messageID,
		WhenUpdated:  time.Now(),
		UserMessage:  message,
		UserFilePath: filePath,
	}

	result, err := messageRepo.Update(messageEntity)
	if err != nil {
		return domain.Message{}, err
	}

	return result, nil
}

func (ms *MessageService) Get(strTopicID string) ([]models.RespMessage, error) {
	topicID, err := uuid.FromString(strTopicID)
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return []models.RespMessage{}, err
	}

	result, err := messageRepo.Get(topicID)
	if err != nil {
		return []models.RespMessage{}, err
	}

	return result, nil
}

func (ms *MessageService) Delete(userID uuid.UUID, strMessageID string) error {
	messageID, err := uuid.FromString(strMessageID)
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return err
	}

	err = messageRepo.Delete(messageID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (ms *MessageService) RatingMessages(strTopicID string, limit, skip uint64) ([]models.RespMessage, error) {
	topicID, err := uuid.FromString(strTopicID)
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return []models.RespMessage{}, err
	}

	result, err := messageRepo.RatingMessagesLike(limit, skip, topicID)
	if err != nil {
		return []models.RespMessage{}, err
	}

	return result, nil
}

func (ms *MessageService) Rating(strLikeBool, strCommentID string, userID uuid.UUID) error {
	commentID, err := uuid.FromString(strCommentID)
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return err
	}

	like, err := strconv.ParseBool(strLikeBool)
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return err
	}

	id, err := uuid.NewV4()
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return err
	}

	ratingEntity := domain.Rating{
		ID:        id,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
		MessageID: commentID,
		UserID:    userID,
	}

	if like {
		if err = messageRepo.CheckDizLike(ratingEntity); err != nil {
			if err = messageRepo.CheckLike(ratingEntity); err != nil {
				err = messageRepo.PostLike(ratingEntity)
				if err != nil {
					log.WithField("component", "service").Debug(err)
					return err
				}
				return nil
			} else {
				err = messageRepo.DeleteLike(ratingEntity)
				if err != nil {
					log.WithField("component", "service").Debug(err)
					return err
				}
				return nil
			}
		} else {
			err = messageRepo.DeleteDizLike(ratingEntity)
			if err != nil {
				log.WithField("component", "service").Debug(err)
				return err
			}

			err = messageRepo.PostLike(ratingEntity)
			if err != nil {
				log.WithField("component", "service").Debug(err)
				return err
			}
			return nil
		}
	} else {
		if err = messageRepo.CheckLike(ratingEntity); err != nil {
			if err = messageRepo.CheckDizLike(ratingEntity); err != nil {
				err = messageRepo.PostDizLike(ratingEntity)
				if err != nil {
					log.WithField("component", "service").Debug(err)
					return err
				}
				return nil
			} else {
				err = messageRepo.DeleteDizLike(ratingEntity)
				if err != nil {
					log.WithField("component", "service").Debug(err)
					return err
				}
				return nil
			}
		} else {
			err = messageRepo.DeleteLike(ratingEntity)
			if err != nil {
				log.WithField("component", "service").Debug(err)
				return err
			}

			err = messageRepo.PostDizLike(ratingEntity)
			if err != nil {
				log.WithField("component", "service").Debug(err)
				return err
			}
			return nil
		}
	}
}
