package model

type TimeAttendanceCheckInRequest struct {
	UserId        string `json:"user_id"`
	ProjectId     string `json:"project_id"`
	ProjectPlace  string `json:"project_place"`
	CheckDateTime string `json:"check_date_time"`
	ImageUrl      string `json:"image_url"`
	CheckStatus   string `json:"check_status"`
	CreatedBy     string `json:"created_by"`
	RefId         string `json:"ref_id"`
}

// UserId        string `db:"user_id" json:"user_id"`
// ProjectId     string `db:"project_id" json:"project_id"`
// ProjectPlace  string `db:"project_place" json:"project_place"`
// CheckDateTime string `db:"check_datetime" json:"check_datetime"`
// ImageUrl       string `db:"image_url"`
// CheckStatus   string `db:"check_status" json:"check_status"`
// CreateBy      string `db:"created_by" json:"created_by"`

// type TimeAttendanceCheckInInput struct {
// 	UserId      string `db:"user_id" json:"user_id"`
// 	ProjectId   int    `db:"check_in_project_id" json:"check_in_project_id"`
// 	Place       string `db:"check_in_place" json:"check_in_place"`
// 	CheckInDate string `db:"check_in_date" json:"check_in_date"`
// }

// {
//	"user_id":"TEST001",
//     "check_in_project_id": 1,
//     "check_in_place": "วิลล่า ราชพฤกษ์-ปิ่นเกล้า",
//     "check_in_datetime": "54321"
// }
