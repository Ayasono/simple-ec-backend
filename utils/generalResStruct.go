package utils

// GeneralResStruct 是通用的响应结构体。
type GeneralResStruct struct {
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}
