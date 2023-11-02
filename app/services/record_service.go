package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/PatipatCha/jeab_ta_service/app/model"
)

type RecordService interface {
	FindMonth(month string) string
	thaiMonthName(month int, monthtype string) string
	RecordListForMobile(input []model.TimeAttendanceReportMobileEntity) []model.TimeAttendanceReportList
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

func RecordListForMobile(input []model.TimeAttendanceReportMobileEntity) []model.TimeAttendanceReportList {
	var list []model.TimeAttendanceReportList

	for index, _ := range input {
		t, _ := time.Parse(time.RFC3339, input[index].Date)

		dayStr := strconv.Itoa(t.Day())

		monthInt := int(t.Month())
		thaiMonth := thaiMonthName(monthInt, "short")

		remark, _ := formatDateInThai(input[index].CheckOutRemark)

		newEntry := model.TimeAttendanceReportList{
			Day:            dayStr,
			Month:          thaiMonth,
			ProjectPlace:   input[index].ProjectPlace,
			CheckInTime:    input[index].CheckInTime,
			CheckOutTime:   input[index].CheckOutTime,
			CheckOutRemark: remark,
			Total:          input[index].TotalHour + " ชั่วโมง " + input[index].TotalMinute + " นาที",
		}

		list = append(list, newEntry)

	}

	return list
}

func thaiMonthName(month int, monthtype string) string {
	var thaiMonthNames = []string{
		"ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.",
		"ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค.",
	}

	if monthtype == "full" {
		thaiMonthNames = []string{
			"มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน",
			"กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม",
		}
	}

	if month >= 1 && month <= 12 {
		return thaiMonthNames[month-1]
	}
	return "Invalid Month"
}

func formatDateInThai(inputDate string) (string, error) {
	// Parse the input date string
	t, err := time.Parse("2006-01-02", inputDate)
	if err != nil {
		return "", err
	}

	// Extract the day and month
	day := t.Day()
	month := t.Month()

	// Get the Thai Buddhist month name
	thaiMonth := thaiMonthName(int(month), "full")

	// Format the result
	result := fmt.Sprintf("(%d %s)", day, thaiMonth)

	return result, nil
}

// func MapReportForMobile(entity []model.TimeAttendanceReportMobileEntity) []model.TimeAttendanceReportList {
// 	res := []model.TimeAttendanceReportList{}
// 	_ = model.Date{}
// 	_ = []model.List{}

// 	_ = model.Date{}
// var listObj = &model.List{}

// for index, _ := range entity {
// 	dateEntity := entity[index].Date
// 	trimmed := strings.Trim(dateEntity, "T00:00:00Z")
// 	t, _ := time.Parse("2006-01-02", trimmed)
// 	print(t.String())

// for i := 0; i < 10; i++ {
// 	book.Categories = append(book.Categories, Category{
// 		Id:   10,
// 		Name: "Vanaraj",
// 	})
// }

// }

// return res
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

// }
