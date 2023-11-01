package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/PatipatCha/jeab_ta_service/app/model"
	"github.com/gofiber/fiber/v2"
)

type AttendanceReportService interface {
	FindMonth(month string) string
}

func FindMonth(month string) model.MonthDefine {
	var output = model.MonthDefine{}
	var defineMo = []model.MonthDefine{}
	content, err := ioutil.ReadFile("./app/json/define/define_month.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(content, &defineMo)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	return output
}

func MapDateModelForMobile(day string, month string) model.Date {
	var obj = model.Date{
		Day:   day,
		Month: month,
	}
	return obj
}

func MapReportForMobile(entity []model.TimeAttendanceReportMobileEntity) fiber.Map {
	_ = []model.TimeAttendanceReportList{}
	_ = model.Date{}
	_ = []model.List{}

	var date = []model.Date{}
	for index, _ := range entity {

		// MapDateModelForMobile()

		date := &model.Date{
			Day:   entity[index].Date,
			Month: entity[index].Date,
		}

		fmt.Println(date)

		// lists := model.List{
		// 	ProjectPlace:   entity[index].ProjectPlace,
		// 	CheckInTime:    entity[index].CheckInTime,
		// 	CheckOutTime:   entity[index].CheckOutTime,
		// 	CheckOutRemark: entity[index].CheckOutRemark,
		// 	Total:          entity[index].TotalHour + " ชั่วโมง " + entity[index].TotalMinute + " นาที",
		// }

		// fmt.Println(lists)

	}

	// res = model.TimeAttendanceReportList{Date: date, Lists: lists}

	// date := model.Date{
	// 	Day: ,

	// }
	// list := model.List{
	// 	ProjectPlace:   "",
	// 	CheckInTime:    "",
	// 	CheckOutTime:   "",
	// 	CheckOutRemark: "",
	// 	Total:          "",
	// }

	// var res = model.TimeAttendanceReportList{
	// 	Date:  date,
	// 	Lists: list,
	// }

	res := fiber.Map{
		"Date": date,
		"List": []model.List{},
	}

	return res
}
