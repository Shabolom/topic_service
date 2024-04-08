package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"service_topic/config"
	"service_topic/internal/domain"
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
		Columns("id", "when_created", "topic_info", "topic_name", "topic_file").
		Values(topic.ID, topic.WhenCreated, topic.TopicInfo, topic.TopicName, topic.TopicFile).
		ToSql()
	if err != nil {
		return err
	}

	_, err = config.Pool.Exec(context.TODO(), sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (tr *TopicRepo) FindOne(what string, value any) (domain.Topic, error) {
	var topic domain.Topic

	sql, args, err := config.Sq.
		Select("*").
		From("topics").
		Where(what+"= $1", value).
		Where("when_deleted IS NULL").
		ToSql()
	if err != nil {
		return domain.Topic{}, err
	}

	row := config.Pool.QueryRow(context.TODO(), sql, args...)

	err = row.Scan(&topic.ID, &topic.WhenCreated, &topic.WhenUpdate, &topic.WhenDeleted,
		&topic.TopicInfo, &topic.TopicName, &topic.TopicFile)
	if err != nil {
		return domain.Topic{}, err
	}

	return topic, nil
}

func (tr *TopicRepo) GetAll() ([]domain.Topic, error) {
	var topics []domain.Topic
	var topic domain.Topic

	sql, args, err := config.Sq.
		Select("id", "topic_info", "topic_name", "topic_file").
		From("topics").
		Where("when_deleted IS NULL").
		ToSql()
	if err != nil {
		return []domain.Topic{}, err
	}

	rows, err := config.Pool.Query(context.TODO(), sql, args...)
	value, err := rows.Values()
	fmt.Println(value, err, 123123123123)

	err2 := rows.Scan(&topic.ID,
		&topic.TopicInfo, &topic.TopicName, &topic.TopicFile)

	if err2 != nil {
		return []domain.Topic{}, err2
	}

	topics = append(topics, topic)

	return topics, nil
}

func (tr *TopicRepo) Update(topic domain.Topic, topicID uuid.UUID) error {

	sql, args, err := config.Sq.
		Update("topics").
		Set("when_update", topic.WhenUpdate).
		Set("topic_info", topic.TopicInfo).
		Set("topic_name", topic.TopicName).
		Set("topic_file", topic.TopicFile).
		Where("id = $5", topicID).
		ToSql()
	if err != nil {
		return err
	}

	_, err = config.Pool.Exec(context.TODO(), sql, args...)
	if err != nil {
		return err
	}

	return nil
}
