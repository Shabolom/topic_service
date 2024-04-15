package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"
	"os"
	"service_topic/config"
	"service_topic/internal/domain"
	"service_topic/internal/models"
	"time"
)

type TopicRepo struct {
}

func NewTopicRepo() *TopicRepo {
	return &TopicRepo{}
}

func (tr *TopicRepo) Create(topic domain.Topic) error {
	_, err := tr.FindOne("topic_name", topic.TopicName)
	if err == nil {
		return errors.New("такой топик уже существует")
	}

	sql, args, err := config.Sq.
		Insert("topics").
		Columns("id", "when_crated", "when_update", "topic_info", "topic_name", "topic_file").
		Values(topic.ID, topic.WhenCreated, topic.WhenUpdate, topic.TopicInfo, topic.TopicName, topic.TopicFile).
		ToSql()
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return err
	}

	_, err = config.Pool.Exec(context.TODO(), sql, args...)
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return err
	}

	return nil
}

func (tr *TopicRepo) FindOne(what string, value any) (domain.Topic, error) {
	var topic domain.Topic

	sql, args, err := config.Sq.
		Select("id", "when_crated", "when_update", "topic_info", "topic_name", "topic_file").
		From("topics").
		Where(what+"= $1", value).
		Where("when_deleted IS NULL").
		ToSql()
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return domain.Topic{}, err
	}

	row := config.Pool.QueryRow(context.TODO(), sql, args...)

	err = row.Scan(&topic.ID, &topic.WhenCreated, &topic.WhenUpdate,
		&topic.TopicInfo, &topic.TopicName, &topic.TopicFile)
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return domain.Topic{}, err
	}

	return topic, nil
}

func (tr *TopicRepo) GetAll(limit, skip uint64) ([]domain.Topic, error) {
	var topics []domain.Topic
	var topic domain.Topic

	sql, args, err := config.Sq.
		Select("id", "when_crated", "when_update", "topic_info", "topic_name", "topic_file").
		From("topics").
		Where("when_deleted IS NULL").
		OrderBy("topic_name", "id").
		Offset(skip).
		Limit(limit).
		ToSql()
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return []domain.Topic{}, err
	}

	rows, err := config.Pool.Query(context.TODO(), sql, args...)
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return []domain.Topic{}, err
	}

	for rows.Next() {
		err2 := rows.Scan(
			&topic.ID,
			&topic.WhenCreated,
			&topic.WhenUpdate,
			&topic.TopicInfo,
			&topic.TopicName,
			&topic.TopicFile)
		if err2 != nil {
			log.WithField("component", "repo").Debug(err2)
			return []domain.Topic{}, err2
		}
		topics = append(topics, topic)
	}
	defer rows.Close()

	return topics, nil
}

func (tr *TopicRepo) Update(topic domain.Topic, topicID uuid.UUID) error {

	oldTopic, err := tr.FindOne("id", topicID)
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return err
	}

	sql, args, err := config.Sq.
		Update("topics").
		Set("when_update", topic.WhenUpdate).
		Set("topic_info", topic.TopicInfo).
		Set("topic_name", topic.TopicName).
		Set("topic_file", topic.TopicFile).
		Where("id = $5", topicID).
		ToSql()
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return err
	}

	_, err = config.Pool.Exec(context.TODO(), sql, args...)
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return err
	}

	err = os.Remove(oldTopic.TopicFile)
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return err
	}

	return nil
}

func (tr *TopicRepo) JoinTopic(usersTopics domain.UsersTopics) error {

	sql, args, err := config.Sq.
		Insert("users_topics").
		Columns("id", "when_crated", "user_id", "topic_id").
		Values(usersTopics.ID, usersTopics.WhenCreated, usersTopics.UserID, usersTopics.TopicID).
		ToSql()
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return err
	}

	_, err = config.Pool.Exec(context.TODO(), sql, args...)
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return err
	}

	return nil
}

func (tr *TopicRepo) DeleteTopic(topicID uuid.UUID) error {

	sql, args, err := config.Sq.
		Update("topics").
		Set("when_deleted", time.Now()).
		Where("id = $2", topicID).
		ToSql()
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return err
	}

	fmt.Println(sql, args)
	_, err = config.Pool.Exec(context.TODO(), sql, args...)
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return err
	}

	err = tr.UsersTopicsClean(topicID)
	if err != nil {
		return err
	}

	return nil
}

func (tr *TopicRepo) UsersTopicsClean(topicID uuid.UUID) error {

	sql, args, err := config.Sq.
		Update("users_topics").
		Set("when_deleted", time.Now()).
		Where("topic_id = $2", topicID).
		ToSql()
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return err
	}

	_, err = config.Pool.Exec(context.TODO(), sql, args...)
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return err
	}

	return nil
}

func (tr *TopicRepo) DeleteUser(topicID, userID uuid.UUID) error {

	sql, args, err := config.Sq.
		Update("users_topics").
		Set("when_deleted", time.Now()).
		Where("topic_id = $2 AND user_id = $3", topicID, userID).
		ToSql()
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return err
	}

	_, err = config.Pool.Exec(context.TODO(), sql, args...)
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return err
	}

	return nil
}

func (tr *TopicRepo) TopicRating(limit, skip uint64) ([]models.TopicRating, error) {
	var topicRating []models.TopicRating
	var topic models.TopicRating

	sql, args, err := config.Sq.
		Select("t.id", "t.when_crated", "t.when_update", "t.topic_info", "t.topic_name", "t.topic_file", "COUNT(ut.topic_id) AS count_users").
		From("users_topics ut").
		Where("ut.when_deleted IS NULL").
		Join("topics t ON ut.topic_id = t.id").
		OrderBy("count_users DESC").
		GroupBy("t.id").
		Limit(limit).
		Offset(skip).
		ToSql()
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return []models.TopicRating{}, err
	}

	rows, err := config.Pool.Query(context.TODO(), sql, args...)
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return []models.TopicRating{}, err
	}

	for rows.Next() {
		err2 := rows.Scan(
			&topic.Topic.ID,
			&topic.Topic.WhenCreated,
			&topic.Topic.WhenUpdate,
			&topic.Topic.TopicInfo,
			&topic.Topic.TopicName,
			&topic.Topic.TopicFile,
			&topic.Users)
		if err2 != nil {
			log.WithField("component", "repo").Debug(err2)
			return []models.TopicRating{}, err2
		}
		topicRating = append(topicRating, topic)
	}

	return topicRating, nil
}
