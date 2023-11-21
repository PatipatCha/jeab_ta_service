package services

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PatipatCha/jeab_ta_service/app/databases"
	"github.com/PatipatCha/jeab_ta_service/app/model"
	"github.com/google/uuid"
)

func StringWithCharset(userId string, projectId string) string {
	// ranText := RandomString(10)
	t := strconv.Itoa(int(time.Now().Unix()))
	// encodedText := base64.StdEncoding.EncodeToString([]byte(t + userId + projectId + ranText))
	encodedText := uuid.New().String() + "-" + base64.StdEncoding.EncodeToString([]byte(t+userId))
	return encodedText
}

// func VaildateUserId(userId string) (bool, error) {
// 	var ta model.TimeAttendanceEntity
// 	db, err := databases.ConnectAccountDB()
// 	if err != nil {
// 		return false, err
// 	}

// 	result := db.Table("users").Where("user_id = ?", userId).Where("status = ?", "active").Where("role = ?", "sg").Scan(&ta)
// 	if result.RowsAffected <= 0 {
// 		return false, nil
// 	}

// 	return true, nil
// }

func VaildateCheckStatus(userId string, refId string, check_status string) (bool, model.TimeAttendanceStatus, string) {
	var ta model.TimeAttendanceStatus
	checkStatus := strings.ToLower(check_status)

	db, err := databases.ConnectTADB()
	if err != nil {
		return false, ta, string(err.Error())
	}

	if checkStatus == "checkin" {
		db.Table("time_attendance").Where("user_id = ?", userId).Order("id DESC").Limit(1).Find(&ta)
		fmt.Println(ta.CheckStatus + " === " + string(checkStatus))
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

	return true, ta, "Passed"
}

func CheckStatus(userId string) (model.TimeAttendanceEntity, error) {
	var ta model.TimeAttendanceEntity

	db, err := databases.ConnectTADB()
	if err != nil {
		return ta, err
	}

	db.Table("time_attendance").Where("user_id = ?", userId).Last(&ta)

	return ta, nil

}
