package model

type TimeAttendanceReportForMobileResponse struct {
	UserId  string                     `json:"user_id"`
	Data    []TimeAttendanceReportList `json:"data"`
	Message string                     `json:"message"`
}

type TimeAttendanceReportList struct {
	Day            string `json:"day"`
	Month          string `json:"month"`
	ProjectPlace   string `json:"project_place"`
	CheckInTime    string `json:"check_in_time"`
	CheckOutTime   string `json:"check_out_time"`
	CheckOutRemark string `json:"check_out_remark"`
	Total          string `json:"total"`
}

// type TimeAttendanceReportList struct {
// 	Date  Date   `json:"date"`
// 	Lists []List `json:"lists"`
// }

// type Date struct {
// 	Day   string `json:"day"`
// 	Month string `json:"month"`
// }

// type List struct {
// 	ProjectPlace   string `json:"project_place"`
// 	CheckInTime    string `json:"check_in_time"`
// 	CheckOutTime   string `json:"check_out_time"`
// 	CheckOutRemark string `json:"check_out_remark"`
// 	Total          string `json:"total"`
// }
