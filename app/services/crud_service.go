package services

import (
	"strings"

	"github.com/PatipatCha/jeab_ta_service/app/databases"
	"github.com/PatipatCha/jeab_ta_service/app/model"
)

type CRUDService interface {
	SaveData(request model.TimeAttendanceCheckInRequest) (model.TimeAttendanceEntity, error)
	GetReportMonth() string
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
