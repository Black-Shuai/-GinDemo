package Models

type LeaveAnswer struct {
	Id int  		`json:"id"`
	Message string 	`json:"message"`
	Next_id int `json:"next_id"`
	Author string 	`json:"author"`
	CreateTime string `json:"create_time"`
}

func (LeaveAnswer)TableName() string  {
	return "tb_leaveanswer"
}
