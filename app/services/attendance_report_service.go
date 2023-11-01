package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/PatipatCha/jeab_ta_service/app/model"
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

func MapReportForMobile(entity []model.TimeAttendanceReportMobileEntity) []model.TimeAttendanceReportList {
	_ = []model.TimeAttendanceReportList{}
	_ = model.Date{}
	_ = []model.List{}

	_ = model.Date{}
	// var listObj = &model.List{}

	for index, _ := range entity {
		dateEntity := entity[index].Date
		trimmed := strings.Trim(dateEntity, "T00:00:00Z")
		t, _ := time.Parse("2006-01-02", trimmed)

		// for i := 0; i < 10; i++ {
		// 	book.Categories = append(book.Categories, Category{
		// 		Id:   10,
		// 		Name: "Vanaraj",
		// 	})
		// }

	}

	return res
	// for indexB, _ := range entity {

	// 	listObj.CheckInTime = ""
	// 	listObj.CheckOutTime = ""
	// 	listObj.CheckOutRemark = entity[index].CheckOutRemark
	// 	listObj.ProjectPlace = ""
	// 	listObj.Total = ""
	// }

	// &date{
	// 	"Date":  entity[index].Date,
	// 	"Month": entity[index].Date,
	// }

	// lists := model.List{
	// 	ProjectPlace:   entity[index].ProjectPlace,
	// 	CheckInTime:    entity[index].CheckInTime,
	// 	CheckOutTime:   entity[index].CheckOutTime,
	// 	CheckOutRemark: entity[index].CheckOutRemark,
	// 	Total:          entity[index].TotalHour + " ชั่วโมง " + entity[index].TotalMinute + " นาที",
	// }

	// fmt.Println(lists)

	// var model = model.TimeAttendanceReportList{
	// 	Date:  dateObj,
	// 	Lists: []model.List{},
	// }

	// dateModel = append(final.Date, model.Date{
	// 	{
	// 		Date:  strconv.Itoa(t.Day()),
	// 		Month: t.Month().String(),
	// 	}
	// })

	// model := fiber.Map{
	// 	"Date": dateObj,
	// 	"List": []model.List{},
	// }

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

	// var res = []model.TimeAttendanceReportList{
	// 	Date:  dateObj,
	// 	Lists: []model.List{},
	// }

	// res := []model.TimeAttendanceReportList{}

}
