package utils

type Response struct {
	Code 	int 		`json:"code"`
	Data 	interface{} `json:"data"`
	Message string 		`json:"message"`
}

func (res *Response) Succ(data interface{}) {
	res.Code = 0
	res.Message = "success"
	res.Data = data
}

func (res *Response) Fail(mes string) {
	res.Code = 1
	res.Message = mes
	res.Data = nil
}

func (res *Response) FailCode(code int, mes string) {
	res.Code = code
	res.Message = mes
	res.Data = nil
}
