package forms

type UpdateBaseProfileData struct {
	FirstName      string `json:"first_name" binding:"required"`
	SecondName     string `json:"second_name" binding:"required"`
	Patronymic     string `json:"patronymic"`
	DateOfBirthday string `json:"date_of_birthday" binding:"required"`
	Time           string `json:"time" binding:"required"`
	Description    string `json:"description" binding:"required"`
}

type UserCommunicationInput struct {
	MessengerId uint   `json:"messenger_id" binding:"required"`
	Login       string `json:"login" binding:"required"`
}

type UserEducationInput struct {
	Institution string `json:"institution" binding:"required"`
	StartYear   int    `json:"start_year" binding:"required"`
	EndYear     int    `json:"end_year" binding:"required"`
	Degree      string `json:"degree" binding:"required"`
}

type UserWorkExperience struct {
	Organization string `json:"organization" binding:"required"`
	StartYear    int    `json:"start_year" binding:"required"`
	EndYear      int    `json:"end_year" binding:"required"`
}

type MentorOtherInfo struct {
	Data string `json:"data" binding:"required"`
}

type UserEmailInput struct {
	Email string `json:"email" binding:"required"`
}

type MentorSpecializationInput struct {
	Specialization string `json:"specialization" binding:"required"`
}
