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
	ImageId       string `db:"image_id"`
	CheckStatus   string `db:"check_status" json:"check_status"`
	CreatedBy     string `db:"created_by" json:"created_by"`
	RefId         string `db:"ref_id" json:"ref_id"`
}
