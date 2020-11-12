package Models

type LeaveMessage struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Leavemessage string `json:"leavemessage"`
	NextId string `json:"next_id"`
	ForAuthor string `json:"for_author"`
	CreateTime string `json:"create_time"`
	LeaveAnswer LeaveAnswer `gorm:"ForeignKey:NextId;"`
}

func (LeaveMessage)TableName() string  {
	return "tb_leavemessage"
}
