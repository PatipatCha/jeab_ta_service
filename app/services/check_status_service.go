package services

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/PatipatCha/jeab_ta_service/app/databases"
	"github.com/PatipatCha/jeab_ta_service/app/model"
)

type CheckStatusService interface {
	StringWithCharset(length int) string
	VaildateCheckStatus(userId string, check_status string) bool
	SaveData(request model.TimeAttendanceCheckInRequest) (model.TimeAttendanceEntity, error)
}

func StringWithCharset(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func VaildateCheckStatus(userId string, refId string, check_status string) bool {
	var ta model.TimeAttendanceEntity
	checkStatus := strings.ToLower(check_status)

	db, err := databases.ConnectDB()
	if err != nil {
		return false
	}

	// res := db.Table("time_attendance").Where("check_status", status).Where("user_id", userId).Where("check_date_time", now.Format("2006-01-02")).Scan(&s)

	if checkStatus == "checkin" {
		db.Table("time_attendance").Where("user_id = ?", userId).Last(&ta)
		println(ta.CheckStatus)
		if ta.CheckStatus == checkStatus {
			return false
		}
	}

	if checkStatus == "checkout" {
		if refId != "" {
			result := db.Table("time_attendance").Where("user_id = ?", userId).Where("ref_id = ?", refId).First(&ta)
			println("checkout" + ta.RefId)
			if result.RowsAffected == 0 {
				log.Fatalf("cannot retrieve : %v\n", result.Error)
				return false
			}
		}

		if ta.RefId == refId {
			println(ta.RefId + "==" + refId)
			return false
		}
	}

	// if result.Error != nil {
	// 	log.Fatalf("cannot retrieve : %v\n", result.Error)
	// 	return false
	// }

	// if ta.CheckStatus == checkStatus {
	// 	return false
	// }

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
		RefId:         request.RefId,
	}

	db, err := databases.ConnectDB()
	if err != nil {
		return entity, err
	}

	err = db.Table("time_attendance").Create(&entity).Scan(&entity).Error

	return entity, err
}
