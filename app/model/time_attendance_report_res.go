package model

type TimeAttendanceReportResponse struct {
	UserId  string                     `json:"user_id"`
	Data    []TimeAttendanceReportList `json:"data"`
	Message string                     `json:"message"`
}
