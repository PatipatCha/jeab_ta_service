package model

import (
	"gorm.io/gorm"
)

type TimeAttendanceEntity struct {
	gorm.Model
	UserId        string `db:"user_id" json:"user_id"`
	ProjectId     string `db:"project_id" json:"project_id"`
	ProjectPlace  string `db:"project_place" json:"project_place"`
	CheckDateTime string `db:"check_date_time" json:"check_date_time"`
	ImageUrl      string `db:"image_url"`
	CheckStatus   string `db:"check_status" json:"check_status"`
	CreatedBy     string `db:"created_by" json:"created_by"`
	RefId         string `db:"ref_id" json:"ref_id"`
}

type TimeAttendanceReportMobileEntity struct {
	Date           string `db:"date" json:"date"`
	ProjectPlace   string `db:"project_place" json:"project_place"`
	CheckInTime    string `db:"check_in_time" json:"check_in_time"`
	CheckOutTime   string `db:"check_out_time" json:"check_out_time"`
	CheckOutRemark string `db:"check_out_remark" json:"check_out_remark"`
	TotalHour      string `db:"total_hour" json:"total_hour"`
	TotalMinute    string `db:"total_minute" json:"total_minute"`
}
