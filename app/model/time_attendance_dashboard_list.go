package model

type TimeAttendanceDashboardList struct {
	UserId        string `json:"user_id"`
	ProjectPlace  string `json:"project_place"`
	CheckInDate   string `json:"check_in_date"`
	CheckInTime   string `json:"check_in_time"`
	CheckInImage  string `json:"check_in_image"`
	CheckOutDate  string `json:"check_out_date"`
	CheckOutTime  string `json:"check_out_time"`
	CheckOutImage string `json:"check_out_image"`
}
