package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/PatipatCha/jeab_ta_service/app/model"
)

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
			Total:          input[index].CheckInDateTime + " ชั่วโมง " + input[index].CheckOutDateTime + " นาที",
		}

		list = append(list, newEntry)

	}

	return list
}

func RecordListJGuard(input []model.TimeAttendanceReportMobileEntity) []model.TimeAttendanceReportList {
	var list []model.TimeAttendanceReportList

	for index, _ := range input {
		t, _ := time.Parse(time.RFC3339, input[index].Date)

		dayStr := strconv.Itoa(t.Day())

		monthInt := int(t.Month())
		thaiMonth := thaiMonthName(monthInt, "short")

		remark, _ := formatDateInThai(input[index].CheckOutRemark)

		var work, _ = CalculateTotalWorkingHours(input[index].CheckInDateTime, input[index].CheckOutDateTime)
		hour, minute, _ := ExtractHoursAndMinutes(work)

		total := hour + " ชั่วโมง " + minute + " นาที"

		newEntry := model.TimeAttendanceReportList{
			Day:            dayStr,
			Month:          thaiMonth,
			ProjectPlace:   input[index].ProjectPlace,
			CheckInTime:    input[index].CheckInTime,
			CheckOutTime:   input[index].CheckOutTime,
			CheckOutRemark: remark,
			Total:          total,
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

func FormatDuration(duration time.Duration) string {
	// Round the duration to the nearest second
	roundedDuration := duration.Round(time.Second)

	hours := int(roundedDuration.Hours())
	minutes := int((roundedDuration - time.Duration(hours)*time.Hour).Minutes())

	fmt.Print(hours)

	var result string
	if hours > 0 {
		result += fmt.Sprintf("%02d:", hours)
	}

	result += fmt.Sprintf("%02d", minutes)

	return result
}

func CalculateTotalWorkingHours(startTime, endTime string) (string, error) {
	// Parse input strings into time.Time objects
	startTimeObj, err := time.Parse(time.RFC3339, startTime)
	if err != nil {
		return "", fmt.Errorf("error parsing start time: %v", err)
	}

	endTimeObj, err := time.Parse(time.RFC3339, endTime)
	if err != nil {
		return "", fmt.Errorf("error parsing end time: %v", err)
	}

	// Check that start time is before end time
	if startTimeObj.After(endTimeObj) {
		return "", fmt.Errorf("start time must be before end time")
	}

	// Calculate working hours
	workingHours := endTimeObj.Sub(startTimeObj)

	// Format workingHours as "HH:MM"
	workingHoursFormatted := fmt.Sprintf("%02d:%02d", int(workingHours.Hours()), int(workingHours.Minutes())%60)

	return workingHoursFormatted, nil
}

func ExtractHoursAndMinutes(workingHoursString string) (string, string, error) {
	parts := strings.Split(workingHoursString, ":")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid format for working hours string")
	}

	hours := parts[0]
	minutes := parts[1]

	return hours, minutes, nil
}
