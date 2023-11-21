package model

type TimeAttendanceCheckInOutRequest struct {
	UserId        string `json:"user_id"`
	ProjectId     string `json:"project_id"`
	ProjectPlace  string `json:"project_place"`
	CheckDateTime string `json:"check_date_time"`
	ImageUrl      string `json:"image_url"`
	CheckStatus   string `json:"check_status"`
	CreatedBy     string `json:"created_by"`
	RefId         string `json:"ref_id"`
}

type TimeAttendanceStatus struct {
	UserId        string `db:"user_id" json:"user_id"`
	ProjectId     string `db:"project_id" json:"project_id"`
	ProjectPlace  string `db:"project_place" json:"project_place"`
	CheckDateTime string `db:"check_date_time" json:"check_date_time"`
	CheckStatus   string `json:"check_status"`
	RefId         string `json:"ref_id"`
}

type TimeAttendanceReportList struct {
	Day            string `json:"day"`
	Month          string `json:"month"`
	ProjectPlace   string `json:"project_place"`
	CheckInTime    string `json:"check_in_time"`
	CheckOutTime   string `json:"check_out_time"`
	CheckOutRemark string `json:"check_out_remark"`
	CheckStatus    string `json:"check_status"`
	Total          string `json:"total"`
}

type TimeAttendanceReportForMobileResponse struct {
	UserId  string                     `json:"user_id"`
	Data    []TimeAttendanceReportList `json:"data"`
	Message string                     `json:"message"`
}

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

type TimeAttendanceDashboardForWebResponse struct {
	UserId  string                        `json:"user_id"`
	Data    []TimeAttendanceDashboardList `json:"data"`
	Message string                        `json:"message"`
}
