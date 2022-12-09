package repository

import (
	"Skipper/pkg/models"
	"gorm.io/gorm"
)

type ReportsPostgres struct {
	db *gorm.DB
}

func NewReportsPostgres(db *gorm.DB) *ReportsPostgres {
	return &ReportsPostgres{db: db}
}

func (r *ReportsPostgres) MakeReport(userId uint, fromUserId uint, reportTheme string, reportText string) error {
	var user models.User
	result := r.db.Where("id=?", userId).First(&user)
	if result.Error != nil {
		return result.Error
	}
	var fromUser models.User
	r.db.Where("id=?", fromUserId).First(&fromUser)
	report := models.Report{
		ReportTheme: reportTheme,
		ReportText:  reportText,
		FromUserId:  fromUser.ID,
		ToUserId:    user.ID,
	}
	fromUser.UserReports = append(fromUser.UserReports, report)
	user.Reports = append(user.UserReports, report)
	//r.db.Save(&user)
	//r.db.Save(&fromUser)
	r.db.Create(&report)
	return nil
}

func (r *ReportsPostgres) GetFromUserReports(userId uint) []models.Report {
	var user models.User
	r.db.Preload("UserReports").Where("id=?", userId).First(&user)
	return user.UserReports
}

func (r *ReportsPostgres) GetToUserReports(userId uint) []models.Report {
	var user models.User
	r.db.Preload("Reports").Where("id=?", userId).First(&user)
	return user.Reports
}
