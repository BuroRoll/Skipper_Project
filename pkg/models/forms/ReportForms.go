package forms

type ReportUser struct {
	UserId      uint   `json:"user_id"`
	ReportTheme string `json:"report_theme"`
	ReportText  string `json:"report_text"`
}
