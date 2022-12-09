package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	FromUserId  uint
	ToUserId    uint
	ReportTheme string
	ReportText  string
}
