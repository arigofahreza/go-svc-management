package models

type Metadata struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`

	TimeExecution float64 `json:"time_execution,omitempty"`
	Size          int     `json:"size,omitempty"`
	TotalPage     int     `json:"total_page,omitempty"`
	TotalData     int     `json:"total_data,omitempty"`
}

type BaseResponse struct {
	Metadata *Metadata   `json:"metadata,omitempty"`
	Data     interface{} `json:"data"`
}

func (response BaseResponse) OverrideMetadata(code int, status string, message string, timeExecution float64) *Metadata {
	return &Metadata{
		Code:          code,
		Status:        status,
		Message:       message,
		TimeExecution: timeExecution,
	}
}

func (response BaseResponse) OverrideMetadataPagination(code int, status string, message string, timeExecution float64, size int, totalPage int, totalData int) *Metadata {
	return &Metadata{
		Code:          code,
		Status:        status,
		Message:       message,
		TimeExecution: timeExecution,
		Size:          size,
		TotalPage:     totalPage,
		TotalData:     totalData,
	}
}
