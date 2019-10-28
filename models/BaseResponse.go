package models

type BaseResponseModel struct {
	//必须的大写开头
	Code int
	Msg  string
	Data interface{} // 这个好像是任意类型的一种体现
}

func SuccessBaseResponse(data interface{}) BaseResponseModel {

	return BaseResponseModel{200, "成功", data}

}
