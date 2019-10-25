package models

type BaseResponseModel struct {
	//必须的大写开头
	Code string
	Msg  string
	Data User
}

type User struct {
	//必须的大写开头
	Id          int
	UserName    string `orm:"size(20);null"`
	Password    string `orm:"size(32);null"`
	Tel         string `orm:"size(11);null"`
	Address     string `orm:"null"`
	PaperNumber int    `orm:"null"`
}
