package service

import (
	"Skipper/pkg/models"
	"Skipper/pkg/models/forms"
	"Skipper/pkg/repository"
	"encoding/json"
)

type CommentsService struct {
	repo repository.Comments
}

func NewCommentsService(repo repository.Comments) *CommentsService {
	return &CommentsService{repo: repo}
}

func (c CommentsService) CreateComment(comment forms.CommentInput) error {
	if comment.Anonymous {
		comment.LessonsCount = nil
		comment.SenderId = nil
	}
	err := c.repo.CreateComment(comment)

	if err != nil {
		return err
	}
	c.CalcRating(comment.RecipienId)
	return nil
}

func (c CommentsService) CreateLessonComment(lessonCommentData forms.CommentInput) error {
	lessonComment := models.LessonComment{
		SenderId:   lessonCommentData.SenderId,
		RecipienId: lessonCommentData.RecipienId,
		Rating:     lessonCommentData.Rating,
	}
	err := c.repo.CreateLessonComment(lessonComment)
	if err != nil {
		return err
	}
	c.CalcRating(lessonComment.RecipienId)
	return nil
}

func (c CommentsService) GetComments(userId uint) (string, error) {
	commentsData, err := c.repo.GetComments(userId)
	if err != nil {
		return "", err
	}
	commentsJson, err := json.Marshal(commentsData)
	return string(commentsJson), nil
}

func (c CommentsService) CalcRating(userId uint) {
	c.repo.CalcRating(userId)
}
