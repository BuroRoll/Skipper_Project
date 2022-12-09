package service

import (
	"Skipper/pkg/models/forms"
	"Skipper/pkg/repository"
)

type ReportsService struct {
	repo repository.Reports
}

func NewReportsService(repo repository.Reports) *ReportsService {
	return &ReportsService{repo: repo}
}

func (r ReportsService) MakeReport(userId uint, reportForm forms.ReportUser) error {
	return r.repo.MakeReport(reportForm.UserId, userId, reportForm.ReportTheme, reportForm.ReportText)
}
