package utils

type Response struct{
	Status bool `json:"status"`
	Erros interface{} `json:"errors"`
	Data interface{}`json:"data"`
}

func GetResponse(data interface{}) *Response{
	return &Response{
		Status: true,
		Erros: nil,
		Data: data,
	}
}

func GetErrorResponse(err interface{}) *Response{
	return &Response{
		Status: false,
		Erros: err,
		Data: nil,
	}
}