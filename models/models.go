package models

type BaseResponseModel struct {
	//必须的大写开头
	Code string
	Msg  string
	Data User
}
