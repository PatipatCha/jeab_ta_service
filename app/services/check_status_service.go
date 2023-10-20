package services

import (
	"log"
	"strings"

	"github.com/PatipatCha/jeab_ta_service/app/databases"
	"github.com/PatipatCha/jeab_ta_service/app/model"
)

type CheckStatusService interface {
	VaildateCheckStatus(userId string, check_status string) bool
	SaveData(request model.TimeAttendanceCheckInRequest) (model.TimeAttendanceEntity, error)
}

func VaildateCheckStatus(userId string, check_status string) bool {
	var ta model.TimeAttendanceEntity
	checkStatus := strings.ToLower(check_status)

	db, err := databases.ConnectDB()
	if err != nil {
		return false
	}

	// res := db.Table("time_attendance").Where("check_status", status).Where("user_id", userId).Where("check_date_time", now.Format("2006-01-02")).Scan(&s)
	result := db.Table("time_attendance").Where("user_id = ?", userId).Last(&ta)
	if result.Error != nil {
		log.Fatalf("cannot retrieve : %v\n", result.Error)
		return false
	}

	if ta.CheckStatus == checkStatus {
		return false
	}

	return true
}

func SaveData(request model.TimeAttendanceCheckInRequest) (model.TimeAttendanceEntity, error) {
	entity := model.TimeAttendanceEntity{
		UserId:        string(request.UserId),
		CheckDateTime: string(request.CheckDateTime),
		ProjectId:     string(request.ProjectId),
		ProjectPlace:  string(request.ProjectPlace),
		CheckStatus:   strings.ToLower(request.CheckStatus),
		CreatedBy:     request.CreatedBy,
		ImageId:       request.ImageId,
	}

	db, err := databases.ConnectDB()
	if err != nil {
		return entity, err
	}

	err = db.Table("time_attendance").Create(&entity).Scan(&entity).Error

	return entity, err
}
