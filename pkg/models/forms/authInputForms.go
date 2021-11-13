package forms

type SignInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenReqBody struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type SignUpUserForm struct {
	Phone      string `json:"phone" binding:"required"`
	Password   string `json:"password" binding:"required"`
	FirstName  string `json:"first_name" binding:"required"`
	SecondName string `json:"second_name" binding:"required"`
}

type SignUpMentorForm struct {
	Phone          string `form:"phone" binding:"required"`
	Password       string `form:"password" binding:"required"`
	FirstName      string `form:"first_name" binding:"required"`
	SecondName     string `form:"second_name" binding:"required"`
	Specialization string `form:"specialization" binding:"required"`
	Description    string `form:"description" binding:"required"`
	Time           string `form:"time" binding:"required"`
}

type SignUpUserToMentorForm struct {
	Specialization string `form:"specialization" binding:"required"`
	Description    string `form:"description" binding:"required"`
	Time           string `form:"time" binding:"required"`
}
