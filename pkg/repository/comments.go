package repository

import (
	"Skipper/pkg/models"
	"Skipper/pkg/models/forms"
	"errors"
	"gorm.io/gorm"
	"math"
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

type CommentData struct {
	gorm.Model
	SenderId             *uint
	RecipienId           uint
	Text                 string
	Rating               uint
	Anonymous            bool
	LessonsCount         *uint
	SenderFirstName      string `json:"sender_first_name"`
	SenderSecondName     string `json:"sender_second_name"`
	SenderProfilePicture string `json:"sender_profile_picture"`
}

func (c CommentsPostgres) GetComments(userId uint) ([]CommentData, error) {
	var comments []CommentData
	result := c.db.
		Select("*").
		Table("comments").
		Joins("LEFT JOIN (SELECT id AS sender_id_data, first_name AS sender_first_name, second_name AS sender_second_name, profile_picture AS sender_profile_picture FROM users) AS sender_data ON sender_data.sender_id_data = sender_id").
		Where("recipien_id=?", userId).
		Find(&comments)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return comments, nil
}

type RatingData struct {
	RecipienId uint    `json:"recipien_id"`
	Rating     float64 `json:"rating"`
}

func (c CommentsPostgres) CalcRating(userId uint) {
	var user models.User
	var classRating RatingData
	var lessonRating RatingData
	c.db.Raw("SELECT recipien_id, avg(rating) AS rating FROM comments WHERE recipien_id = ? group by recipien_id;", userId).
		Find(&classRating)
	c.db.Raw("SELECT recipien_id, avg(rating) AS rating FROM lesson_comments WHERE recipien_id = ? group by recipien_id;", userId).
		Find(&lessonRating)
	c.db.Find(&user, userId)
	rating := 0.0
	if lessonRating.Rating != 0 {
		rating = math.Round((classRating.Rating*0.8+lessonRating.Rating*0.2)*100) / 100
	} else {
		rating = math.Round(classRating.Rating*100) / 100
	}
	user.Rating = rating
	c.db.Save(&user)
}

func (c CommentsPostgres) CreateLessonComment(lessonComment models.LessonComment) error {
	result := c.db.Create(&lessonComment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
