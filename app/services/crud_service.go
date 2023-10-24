package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	"github.com/PatipatCha/jeab_ta_service/app/databases"
	"github.com/PatipatCha/jeab_ta_service/app/model"
)

type CRUDService interface {
	SaveData(request model.TimeAttendanceCheckInRequest) (model.TimeAttendanceEntity, error)
	GetReportByMonth() model.TimeAttendanceReportList
	GetReportNow() model.TimeAttendanceReportList
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

func GetReportByMonth(month string) model.TimeAttendanceReportList {
	var ta_report = model.TimeAttendanceReportList{}

	return ta_report

}

func GetReportNow() model.TimeAttendanceReportList {
	var ta_report = model.TimeAttendanceReportList{}
	content, err := ioutil.ReadFile("./app/json/record_mockup_c.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(content, &ta_report)
	log.Printf("origin: %s\n", ta_report.Data)
	return ta_report
}

// func GetReport(c *fiber.Ctx) error {
// 	var ta []model.TimeAttendanceEntity
// 	db, err := databases.ConnectDB()
// 	if err != nil {
// 		return err
// 	}

// 	// userId := c.Params("userId")
// 	res := db.Table("time_attendance").Find(&ta).Error
// 	if err := res; err != nil {
// 		return err
// 	}

// 	println()

// 	return c.JSON(ta)
// }
