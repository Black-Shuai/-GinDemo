package Models

type LeaveAnswer struct {
	Id int `json:"id"`
	Message string `json:"message"`
	CreateTime string `json:"create_time"`
}

func (LeaveAnswer)TableName() string  {
	return "tb_leaveanswer"
}
