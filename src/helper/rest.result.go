package helper

import "github.com/gin-gonic/gin"

type response struct {
	Meta meta  `json:"meta"`
	Data gin.H `json:"data"`
}

type meta struct {
	Message interface{} `json:"message"`
	Status  string      `json:"status"`
}

func NewRestResult() *response {
	return &response{}
}

func (b *response) SetMeta(status string, msg interface{}) *response {
	b.Meta = meta{
		Message: msg,
		Status:  status,
	}

	return b
}

func (b *response) SetData(data gin.H) *response {
	b.Data = data
	return b
}

func ResponseWithJson(status string, msg interface{}, data gin.H) *response {
	response := NewRestResult()
	response.SetMeta(status, msg)
	response.SetData(data)
	return response
}

func ResponseValidationError(status string, err error) *response {
	response := NewRestResult()
	response.SetMeta(status, FormatValidationError(err))
	return response
}
