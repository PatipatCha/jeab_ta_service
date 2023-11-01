package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/PatipatCha/jeab_ta_service/app/model"
	"github.com/PatipatCha/jeab_ta_service/app/services"
	"github.com/gofiber/fiber/v2"
)

func GetReportWeb(c *fiber.Ctx) error {
	var taReportList = []model.TimeAttendanceReportList{}
	// var ta []model.TimeAttendanceEntity
	// userId := c.Params("userId")
	_ = c.Query("month")
	userId := c.Query("user_id")
	resValidate, _ := services.VaildateUserId(userId)
	if !resValidate {
		output := model.TimeAttendanceResponse{
			UserId:  userId,
			Data:    taReportList,
			Message: os.Getenv("VAILD_USERID_NOT_FOUND"),
		}
		return c.JSON(output)
	}

	findUserId := c.Query("find_user_id")
	data, _ := services.GetReportForWeb(findUserId)

	var res = model.TimeAttendanceDashboardForWebResponse{
		UserId:  userId,
		Data:    data,
		Message: "Attendance Report List",
	}

	return c.JSON(res)
}

func GetReportMobile(c *fiber.Ctx) error {
	var taReportList = []model.TimeAttendanceReportList{}
	userId := c.Query("user_id")
	resValidate, _ := services.VaildateUserId(userId)
	if !resValidate {
		output := model.TimeAttendanceResponse{
			UserId:  userId,
			Data:    taReportList,
			Message: os.Getenv("VAILD_USERID_NOT_FOUND"),
		}
		return c.JSON(output)
	}

	month := c.Query("month")
	_, data, msg := services.GetReportForMobile(userId, month)

	resData := services.MapReportForMobile(data)

	var res = fiber.Map{
		"UserId":  userId,
		"Data":    resData,
		"Message": msg,
	}

	return c.JSON(res)
}

func GetReportMockUp(c *fiber.Ctx) error {
	userId := c.Query("user_id")
	var ta_report = []model.TimeAttendanceReportList{}
	content, err := ioutil.ReadFile("./app/json/record_mockup_test_mobile.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(content, &ta_report)

	// return ta_report

	// var taReportList = []model.TimeAttendanceReportList{}

	var res = model.TimeAttendanceReportForMobileResponse{
		UserId:  userId,
		Data:    ta_report,
		Message: "mockup",
	}

	return c.JSON(res)
}

// for _, taParam := range ta {
// 	// fmt.Println(integ, " = ", spell.CheckInAT)
// 	taModel := model.TimeAttendance{CheckInDate: taParam.CheckInAT, CheckInPlace: taParam.CheckInPlace, CheckInTime: taParam.CheckInAT, CheckOutTime: taParam.CheckOutAt, CheckOutPlace: taParam.CheckOutPlace}
// }

// a := model.TimeAttendanceResponse{user_id: "", Message: "", list: taModel}

// bytes, _ := json.Marshal(ta)
// fmt.Println(string())
