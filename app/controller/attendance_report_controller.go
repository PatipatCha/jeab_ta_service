package controller

import (
	"github.com/PatipatCha/jeab_ta_service/app/databases"
	"github.com/PatipatCha/jeab_ta_service/app/model"
	"github.com/gofiber/fiber/v2"
)

func GetReport(c *fiber.Ctx) error {
	var ta []model.TimeAttendanceEntity
	db, err := databases.ConnectTADB()
	if err != nil {
		return err
	}

	// userId := c.Params("userId")
	res := db.Table("time_attendance").Find(&ta).Error
	if err := res; err != nil {
		return err
	}

	// for _, taParam := range ta {
	// 	// fmt.Println(integ, " = ", spell.CheckInAT)
	// 	taModel := model.TimeAttendance{CheckInDate: taParam.CheckInAT, CheckInPlace: taParam.CheckInPlace, CheckInTime: taParam.CheckInAT, CheckOutTime: taParam.CheckOutAt, CheckOutPlace: taParam.CheckOutPlace}
	// }

	// a := model.TimeAttendanceResponse{user_id: "", Message: "", list: taModel}

	// bytes, _ := json.Marshal(ta)
	// fmt.Println(string())

	return res
}
