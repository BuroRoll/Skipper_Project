package repository

import (
	"Skipper/pkg/models"
	"Skipper/pkg/models/forms"
	"errors"
	"gorm.io/gorm"
)

type CommentsPostgres struct {
	db *gorm.DB
}

func NewCommentsPostgres(db *gorm.DB) *CommentsPostgres {
	return &CommentsPostgres{db: db}
}

func (c CommentsPostgres) CreateComment(commInput forms.CommentInput) error {
	comment := models.Comment{
		SenderId:     commInput.SenderId,
		RecipienId:   commInput.RecipienId,
		Text:         commInput.Text,
		Rating:       commInput.Rating,
		Anonymous:    commInput.Anonymous,
		LessonsCount: commInput.LessonsCount,
	}

	result := c.db.Create(&comment)
	if errors.Is(result.Error, gorm.ErrRegistered) {
		return gorm.ErrRegistered
	}
	return nil
}

func (c CommentsPostgres) GetComments(userId uint) ([]models.Comment, error) {
	var comments []models.Comment
	result := c.db.Where("recipien_id=?", userId).Find(&comments)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return comments, nil
}
