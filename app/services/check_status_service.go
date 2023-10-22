package services

import (
	"encoding/base64"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PatipatCha/jeab_ta_service/app/databases"
	"github.com/PatipatCha/jeab_ta_service/app/model"
)

type CheckStatusService interface {
	RandomString(number int) string
	StringWithCharset(length int) string
	VaildateUserId(userId string) bool // Call Api User Microservices : Phase 2
	VaildateCheckStatus(userId string, check_status string) bool
	SaveData(request model.TimeAttendanceCheckInRequest) (model.TimeAttendanceEntity, error)
}

func RandomString(number int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	randomString := make([]byte, number)
	for i := range randomString {
		randomString[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(randomString)
}

func StringWithCharset(userId string, projectId string) string {
	ranText := RandomString(10)
	t := strconv.Itoa(int(time.Now().Unix()))
	encodedText := base64.StdEncoding.EncodeToString([]byte(t + userId + projectId + ranText))
	return encodedText
}

func VaildateUserId(userId string) (bool, error) {
	var ta model.TimeAttendanceEntity
	db, err := databases.ConnectAccountDB()
	if err != nil {
		return false, err
	}

	result := db.Table("users").Where("user_id = ?", userId).Where("status = ?", "active").Where("role = ?", "sg").Scan(&ta)
	if result.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}

func VaildateCheckStatus(userId string, refId string, check_status string) (bool, model.TimeAttendanceEntity, string) {
	var ta model.TimeAttendanceEntity
	checkStatus := strings.ToLower(check_status)

	db, err := databases.ConnectTADB()
	if err != nil {
		return false, ta, string(err.Error())
	}

	if checkStatus == "checkin" {
		db.Table("time_attendance").Where("user_id = ?", userId).Last(&ta)
		// fmt.Println(ta.CheckStatus + " === " + string(checkStatus))
		if ta.CheckStatus == checkStatus {
			return false, ta, os.Getenv("VAILD_CHECKIN_ERROR")
		}
	}

	if checkStatus == "checkout" {
		findCheckIn := db.Table("time_attendance").Where("user_id = ?", userId).Where("ref_id = ?", refId).Where("check_status = ?", "checkin").Last(&ta)
		if findCheckIn.RowsAffected == 0 {
			return false, ta, "Ref ID Not Found"
		}

		var result model.TimeAttendanceEntity
		db.Table("time_attendance").Where("user_id = ?", userId).Where("ref_id = ?", refId).Where("check_status = ?", "checkout").Last(&result)
		if result.RefId == refId {
			return false, ta, os.Getenv("VAILD_CHECKOUT_ERROR")
		}

	}

	// if result.Error != nil {
	// 	log.Fatalf("cannot retrieve : %v\n", result.Error)
	// 	return false
	// }

	// if ta.CheckStatus == checkStatus {
	// 	return false
	// }

	return true, ta, "Passed"
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

	db, err := databases.ConnectTADB()
	if err != nil {
		return entity, err
	}

	err = db.Table("time_attendance").Create(&entity).Scan(&entity).Error

	// err = db.Table("time_attendance").Find("user_id").Error

	return entity, err
}

// res := db.Table("time_attendance").Where("check_status", status).Where("user_id", userId).Where("check_date_time", now.Format("2006-01-02")).Scan(&s)
