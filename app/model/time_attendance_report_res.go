package model

type TimeAttendanceReportForMobileResponse struct {
	UserId  string                     `json:"user_id"`
	Data    []TimeAttendanceReportList `json:"data"`
	Message string                     `json:"message"`
}

type TimeAttendanceDashboardForWebResponse struct {
	UserId  string                        `json:"user_id"`
	Data    []TimeAttendanceDashboardList `json:"data"`
	Message string                        `json:"message"`
}
