package helper


type response struct {
	Meta  meta        `json:"meta"`
	Data  interface{} `json:"data"`
	Errors interface{} `json:"error"`
}

type meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}


func NewRestResult() *response {
	return &response{}
}

func (b *response) SetMeta(code int, status string, msg string) *response{
	b.Meta = meta{
		Message: msg,
		Code: code,
		Status: status,
	}

	return b
}

func (b *response) SetData(data interface{})  *response {
	b.Data = data
	return b
}

func (b *response) SetErrors(err interface{}) *response {
	b.Errors = err
	return b
}



