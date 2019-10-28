package models

type Model struct {
	//必须的大写开头
	Code string
	Msg  string
	Data interface{} // 这个好像是任意类型的一种体现
}
