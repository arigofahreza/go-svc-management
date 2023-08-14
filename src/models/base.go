package models

type Metadata struct {
	Code          int     `json:"code"`
	Status        string  `json:"status"`
	Message       string  `json:"message"`
	TimeExecution float64 `json:"time_execution"`
}
type BaseResponse struct {
	Metadata *Metadata   `json:"metadata"`
	Data     interface{} `json:"data"`
}

func OverrideMetadata(code int, status string, message string, timeExecution float64) *Metadata {
	return &Metadata{
		Code:          code,
		Status:        status,
		Message:       message,
		TimeExecution: timeExecution,
	}
}
