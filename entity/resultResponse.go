package entity

type SuccessResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type ResultData struct {
	Id int			 `json:"id"`
	Firstname string `json:"firstname"`
	Lastname string  `json:"lastname"`
}
