package service

import (
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
