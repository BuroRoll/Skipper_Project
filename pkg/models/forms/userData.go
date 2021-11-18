package forms

type UpdateBaseProfileData struct {
	FirstName      string `json:"first_name" binding:"required"`
	SecondName     string `json:"second_name" binding:"required"`
	Patronymic     string `json:"patronymic"`
	DateOfBirthday string `json:"date_of_birthday" binding:"required"`
	Time           string `json:"time" binding:"required"`
}
