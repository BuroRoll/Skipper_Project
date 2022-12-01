package forms

type UpdateBaseProfileData struct {
	FirstName      string `json:"first_name"`
	SecondName     string `json:"second_name"`
	Patronymic     string `json:"patronymic"`
	DateOfBirthday string `json:"date_of_birthday"`
	Time           string `json:"time"`
	Description    string `json:"description"`
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
	StartYear    string `json:"start_year" binding:"required"`
	EndYear      string `json:"end_year" binding:"required"`
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

type PasswordChangeInput struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type AddUserToFavourite struct {
	UserId uint   `json:"user_id"`
	Status string `json:"status"`
}

type DeleteFromFavourite struct {
	UserId uint   `json:"user_id"`
	Status string `json:"status"`
}
