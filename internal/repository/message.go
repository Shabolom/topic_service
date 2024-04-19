package repository

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"
	"service_topic/config"
	"service_topic/internal/domain"
	"service_topic/internal/models"
	"service_topic/internal/tools"
	"strings"
	"time"
)

var CommentsHash = map[uuid.UUID][]models.RespMessage{}

type MessageRepo struct {
}

func NewMessageRepo() *MessageRepo {
	return &MessageRepo{}
}

func (mr *MessageRepo) Post(message domain.Message) (domain.Message, error) {

	sql, args, err := config.Sq.
		Insert("comments").
		Columns("id", "when_crated", "when_update", "user_massage", "user_file_path", "user_id", "topic_id").
		Values(
			&message.ID,
			&message.WhenCreated,
			&message.WhenUpdated,
			&message.UserMessage,
			&message.UserFilePath,
			&message.UserID,
			&message.TopicID).
		ToSql()
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return domain.Message{}, err
	}

	_, err = config.Pool.Exec(context.TODO(), sql, args...)
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return domain.Message{}, err
	}

	err = mr.LikeDizLikeCounter(message.ID)
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return domain.Message{}, err
	}

	return message, nil
}

func (mr *MessageRepo) Update(message domain.Message) (domain.Message, error) {
	sql, args, err := config.Sq.
		Update("comments").
		Set("when_update", message.WhenUpdated).
		Set("user_massage", message.UserMessage).
		Set("user_file_path", message.UserFilePath).
		Where("id = $4", message.ID).
		ToSql()
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return domain.Message{}, err
	}

	_, err = config.Pool.Exec(context.TODO(), sql, args...)
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return domain.Message{}, err
	}

	return message, nil
}

//func (mr *MessageRepo) Get(topicID uuid.UUID, limit, skip uint64) ([]models.RespMessage, error) {
//	var message models.RespMessage
//	var messages []models.RespMessage
//	var pathFile string
//
//	commentsID, err := mr.SelectCommentsID(topicID)
//	if err != nil {
//		return []models.RespMessage{}, err
//	}
//
//	for _, commentID := range commentsID {
//		err = mr.UpdateLikeDizLikeCount(commentID)
//		if err != nil {
//			return []models.RespMessage{}, err
//		}
//	}
//
//	sql, args, err := config.Sq.
//		Select(
//			"u.id",
//			"c.id",
//			"u.user_name",
//			"c.user_massage",
//			"c.when_crated",
//			"c.when_update",
//			"c.user_file_path",
//			"ld.likes",
//			"ld.diz_likes").
//		From("users u").
//		Join("comments c ON u.id = c.user_id").
//		Where("c.topic_id = $1 AND c.when_deleted IS NULL", topicID).
//		Join("like_dizlike_count ld ON c.id = ld.comment_id").
//		GroupBy("u.id", "c.id", "ld.likes", "ld.diz_likes").
//		OrderBy("c.when_crated ASC").
//		Offset(skip).
//		Limit(limit).
//		ToSql()
//	if err != nil {
//		log.WithField("component", "repo").Debug(err)
//		return []models.RespMessage{}, err
//	}
//
//	rows, err := config.Pool.Query(context.TODO(), sql, args...)
//	if err != nil {
//		log.WithField("component", "repo").Debug(err)
//		return []models.RespMessage{}, err
//	}
//
//	fmt.Println(sql, args)
//	for rows.Next() {
//
//		err2 := rows.Scan(
//			&message.UserID,
//			&message.MessageID,
//			&message.UserLogin,
//			&message.Message,
//			&message.WhenCreated,
//			&message.WhenUpdate,
//			&pathFile,
//			&message.Like,
//			&message.DizLike,
//		)
//
//		if err2 != nil {
//			log.WithField("component", "repo").Debug(err2)
//			return []models.RespMessage{}, err
//		}
//
//		message.PathToFiles = strings.Split(pathFile, "(space)")
//
//		messages = append(messages, message)
//	}

//	return messages, nil
//}

func (mr *MessageRepo) Get(topicID uuid.UUID, limit, skip uint64) ([]models.RespMessage, error) {
	var message models.RespMessage
	var messages []models.RespMessage
	var pathFile string
	var id uuid.UUID
	var ids []uuid.UUID

	sql, args, err := config.Sq.
		Select("id").
		From("comments c").
		Where("c.topic_id = $1 AND c.when_deleted IS NULL", topicID).
		GroupBy("c.id").
		OrderBy("c.when_crated ASC").
		Offset(skip).
		Limit(limit).
		ToSql()
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return []models.RespMessage{}, err
	}

	rows, err := config.Pool.Query(context.TODO(), sql, args...)
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return []models.RespMessage{}, err
	}

	for rows.Next() {

		err2 := rows.Scan(
			&id,
		)

		if err2 != nil {
			log.WithField("component", "repo").Debug(err2)
			return []models.RespMessage{}, err
		}

		ids = append(ids, id)
	}

	sql, args, err = config.Sq.
		Select(
			"u.id",
			"c.id",
			"u.user_name",
			"c.user_massage",
			"c.when_crated",
			"c.when_update",
			"c.user_file_path",
			"ld.likes",
			"ld.diz_likes").
		From("users u").
		Join("comments c ON u.id = c.user_id").
		Where("c.id = ANY ($1) AND c.when_deleted IS NULL", ids).
		Join("like_dizlike_count ld ON c.id = ld.comment_id").
		GroupBy("u.id", "c.id", "ld.likes", "ld.diz_likes").
		OrderBy("c.when_crated ASC").
		Offset(skip).
		Limit(limit).
		ToSql()
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return []models.RespMessage{}, err
	}

	rows, err = config.Pool.Query(context.TODO(), sql, args...)
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return []models.RespMessage{}, err
	}

	for rows.Next() {

		fmt.Println(sql, args)

		err2 := rows.Scan(
			&message.UserID,
			&message.MessageID,
			&message.UserLogin,
			&message.Message,
			&message.WhenCreated,
			&message.WhenUpdate,
			&pathFile,
			&message.Like,
			&message.DizLike,
		)

		if err2 != nil {
			log.WithField("component", "repo").Debug(err2)
			return []models.RespMessage{}, err
		}

		message.PathToFiles = strings.Split(pathFile, "(space)")

		messages = append(messages, message)
	}

	return messages, nil
}

func (mr *MessageRepo) Delete(messageID, userID uuid.UUID) error {

	sql, args, err := config.Sq.
		Update("comments").
		Set("when_deleted", time.Now()).
		Where("(user_id = $1 OR $2 = $3)", userID, userID, uuid.Nil).
		Where("comment_id = $4", messageID).
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

func (mr *MessageRepo) RatingMessagesLike(limit, skip uint64, topicID uuid.UUID) ([]models.RespMessage, error) {
	var message models.RespMessage
	var messages []models.RespMessage
	var path string

	sql, args, err := config.Sq.
		Select(
			"u.id",
			"c.id",
			"u.user_name",
			"c.user_massage",
			"c.when_crated",
			"c.when_update",
			"c.user_file_path",
			"COUNT(l.user_id) as like",
			"COUNT(dl.user_id) as diz_like").
		From("comments c").
		Where("c.topic_id = $1 AND when_deleted IS NULL", topicID).
		Join("users u ON u.id = c.user_id").
		Join("likes l ON l.comment_id = c.id").
		Join("diz_likes dl ON dl.comment_id = c.id").
		Limit(limit).
		Offset(skip).
		OrderBy("like ASC").
		ToSql()
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return []models.RespMessage{}, err
	}

	rows, err := config.Pool.Query(context.TODO(), sql, args...)
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return []models.RespMessage{}, err
	}

	for rows.Next() {
		err2 := rows.Scan(
			&message.UserID,
			&message.MessageID,
			&message.UserLogin,
			&message.Message,
			&message.WhenCreated,
			&message.WhenUpdate,
			&path,
			&message.Like,
			&message.DizLike)
		if err2 != nil {
			log.WithField("component", "repo").Debug(err2)
			return []models.RespMessage{}, err2
		}

		message.PathToFiles = strings.Split(path, "(space)")

		messages = append(messages, message)
	}

	return messages, nil
}

func (mr *MessageRepo) CheckSubscribe(userID, topicID uuid.UUID) error {
	var id uuid.UUID

	sql, args, err := config.Sq.
		Select("id").
		From("users_topics").
		Where("topic_id = $1", topicID).
		Where("user_id = $2 AND when_deleted IS NULL", userID).
		ToSql()
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return err
	}

	row := config.Pool.QueryRow(context.TODO(), sql, args...)
	err2 := row.Scan(&id)

	if err2 != nil {
		log.WithField("component", "repo").Debug(err2)
		return err2
	}

	return nil
}

func (mr *MessageRepo) LikeDizLikeCounter(messageID uuid.UUID) error {
	id, err := uuid.NewV4()
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return err
	}

	sql, args, err := config.Sq.
		Insert("like_dizlike_count").
		Columns("id", "comment_id", "likes", "diz_likes").
		Values(id, messageID, 0, 0).
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

func (mr *MessageRepo) PostLike(rating domain.Rating) error {
	sql, args, err := config.Sq.
		Insert("likes").
		Columns("id", "when_crated", "when_update", "comment_id", "user_id").
		Values(rating.ID, rating.CreatedAt, rating.UpdateAt, rating.MessageID, rating.UserID).
		ToSql()

	_, err = config.Pool.Exec(context.TODO(), sql, args...)
	if err != nil {
		log.WithField("component", "repo PostLike").Debug(err)
		return err
	}

	return nil
}

func (mr *MessageRepo) PostDizLike(rating domain.Rating) error {
	sql, args, err := config.Sq.
		Insert("diz_likes").
		Columns("id", "when_crated", "when_update", "comment_id", "user_id").
		Values(rating.ID, rating.CreatedAt, rating.UpdateAt, rating.MessageID, rating.UserID).
		ToSql()

	_, err = config.Pool.Exec(context.TODO(), sql, args...)
	if err != nil {
		log.WithField("component", "repo PostDizLike").Debug(err)
		return err
	}

	return nil
}

func (mr *MessageRepo) DeleteDizLike(rating domain.Rating) error {
	sql, args, err := config.Sq.
		Delete("diz_likes").
		Where("comment_id = $1", rating.MessageID).
		Where("user_id = $2", rating.UserID).
		ToSql()

	_, err = config.Pool.Exec(context.TODO(), sql, args...)
	if err != nil {
		tools.INFO.WithField("component", "repo DeleteDizLike").Info(err)
		return err
	}

	return nil
}

func (mr *MessageRepo) DeleteLike(rating domain.Rating) error {
	sql, args, err := config.Sq.
		Delete("likes").
		Where("comment_id = $1", rating.MessageID).
		Where("user_id = $2", rating.UserID).
		ToSql()

	_, err = config.Pool.Exec(context.TODO(), sql, args...)
	if err != nil {
		tools.INFO.WithField("component", "repo DeleteLike").Info(err)
		return err
	}

	return nil
}

func (mr *MessageRepo) CheckLike(rating domain.Rating) error {
	var id uuid.UUID
	sql, args, err := config.Sq.
		Select("id").
		From("likes").
		Where("comment_id = $1", rating.MessageID).
		Where("user_id = $2 AND when_deleted IS NULL", rating.UserID).
		ToSql()

	row := config.Pool.QueryRow(context.TODO(), sql, args...)

	err = row.Scan(&id)
	if err != nil {
		tools.INFO.WithField("component CheckLike", "repo CheckLike").Info(err)
		return err
	}

	return nil
}

func (mr *MessageRepo) CheckDizLike(rating domain.Rating) error {
	var id uuid.UUID
	sql, args, err := config.Sq.
		Select("id").
		From("diz_likes").
		Where("comment_id = $1", rating.MessageID).
		Where("user_id = $2 AND when_deleted IS NULL", rating.UserID).
		ToSql()

	row := config.Pool.QueryRow(context.TODO(), sql, args...)

	err = row.Scan(&id)
	if err != nil {
		tools.INFO.WithField("component", "repo CheckDizLike").Info(err)
		return err
	}

	return nil
}

func (mr *MessageRepo) UpdateLikeDizLikeCount(commentID uuid.UUID) error {

	likeCount, err := mr.LikeCount(commentID)
	if err != nil {
		return err
	}

	DizLikeCount, err := mr.DizLikeCount(commentID)
	if err != nil {
		return err
	}

	sql, args, err := config.Sq.
		Update("like_dizlike_count").
		Set("likes", likeCount).
		Set("diz_likes", DizLikeCount).
		Where("comment_id = $3", commentID).
		ToSql()

	_, err = config.Pool.Exec(context.TODO(), sql, args...)
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return err
	}

	return nil
}

func (mr *MessageRepo) LikeCount(commentID uuid.UUID) (int, error) {
	var count int
	sql, args, err := config.Sq.
		Select("COUNT(id)").
		From("likes").
		Where("comment_id = $1", commentID).
		ToSql()

	row := config.Pool.QueryRow(context.TODO(), sql, args...)

	err = row.Scan(&count)
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return 0, err
	}

	return count, nil
}

func (mr *MessageRepo) DizLikeCount(commentID uuid.UUID) (int, error) {
	var count int
	sql, args, err := config.Sq.
		Select("COUNT(id)").
		From("diz_likes").
		Where("comment_id = $1", commentID).
		ToSql()

	row := config.Pool.QueryRow(context.TODO(), sql, args...)

	err = row.Scan(&count)
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return 0, err
	}

	return count, nil
}

func (mr *MessageRepo) SelectCommentsID(topicID uuid.UUID) ([]uuid.UUID, error) {
	var messageID uuid.UUID
	var messagesID []uuid.UUID

	sql, args, err := config.Sq.
		Select("id").
		From("comments").
		Where("topic_id = $1", topicID).
		ToSql()

	rows, err := config.Pool.Query(context.TODO(), sql, args...)
	if err != nil {
		log.WithField("component", "repo").Debug(err)
		return []uuid.UUID{}, err
	}

	for rows.Next() {
		err2 := rows.Scan(&messageID)
		if err2 != nil {
			log.WithField("component", "repo").Debug(err)
			return []uuid.UUID{}, err2
		}
		messagesID = append(messagesID, messageID)
	}

	return messagesID, nil
}

func HashComments() {
	var messages []models.RespMessage
	var message models.RespMessage
	var pathFile string
	var topicID uuid.UUID

	clear(CommentsHash)

	sql, args, err := config.Sq.
		Select(
			"u.id",
			"c.id",
			"u.user_name",
			"c.user_massage",
			"c.when_crated",
			"c.when_update",
			"c.user_file_path",
			"ld.likes",
			"ld.diz_likes",
			"c.topic_id",
		).
		From("users u").
		Join("comments c ON u.id = c.user_id").
		Where("c.when_deleted IS NULL").
		Join("like_dizlike_count ld ON c.id = ld.comment_id").
		GroupBy("u.id", "c.id", "ld.likes", "ld.diz_likes", "c.topic_id").
		OrderBy("c.when_crated ASC").
		ToSql()
	if err != nil {
		log.WithField("component", "repo-cron").Debug(err)
	}

	rows, err := config.Pool.Query(context.TODO(), sql, args...)
	if err != nil {
		log.WithField("component", "repo-cron").Debug(err)
	}

	for rows.Next() {

		err2 := rows.Scan(
			&message.UserID,
			&message.MessageID,
			&message.UserLogin,
			&message.Message,
			&message.WhenCreated,
			&message.WhenUpdate,
			&pathFile,
			&message.Like,
			&message.DizLike,
			&topicID,
		)

		if err2 != nil {
			log.WithField("component", "repo-cron").Debug(err2)
		}

		message.PathToFiles = strings.Split(pathFile, "(space)")

		messages = append(messages, message)

		CommentsHash[topicID] = append(CommentsHash[topicID], message)
	}

	tools.INFO.WithField("component", "repo-cron").Info("хэш по топикам обновлен")
}
