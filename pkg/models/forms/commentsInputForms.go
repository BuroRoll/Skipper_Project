package forms

type CommentInput struct {
	SenderId     *uint  `json:"sender_id"`
	RecipienId   uint   `json:"recipien_id"`
	Text         string `json:"text"`
	Rating       uint   `json:"rating"`
	Anonymous    bool   `json:"anonymous"`
	LessonsCount *uint  `json:"lessons_count"`
}
