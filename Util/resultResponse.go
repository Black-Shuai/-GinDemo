package Util

import (
	"GinDemo/Dao"
	"GinDemo/Models"
)

type ResultResponse struct {
	Code int
	Message string
	Data interface{}
}

func (this *ResultResponse)ErrResult()(resResponse ResultResponse)  {
	resResponse.Code=0
	resResponse.Message="error"
	return
}
func (this *ResultResponse)SuccessResult(user Models.User)(reResponse ResultResponse) {
	reResponse.Code=1
	reResponse.Message="success"
	reResponse.Data=user
	return
}
func (this *ResultResponse)SuccessResultList(user []Models.User)(reResponse ResultResponse) {
	reResponse.Code=1
	reResponse.Message="success"
	reResponse.Data=user
	return
}
func (this *ResultResponse)SuccessResultList1(user Dao.Resultq)(reResponse ResultResponse) {
	reResponse.Code=1
	reResponse.Message="success"
	reResponse.Data=user
	return
}
