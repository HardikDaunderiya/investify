package types

type BaseHttpResponse struct {
	Status     string      `json:"status"`
	Data       interface{} `json:"data"`
	StatusCode int         `json:"statusCode"`
}
type BaseErrorResponse struct {
	Status     string      `json:"status"`
	Message    interface{} `json:"error"`
	StatusCode int         `json:"statusCode"`
	Extra      string      `json:"extra"`
}

func GenerateResponse(data interface{}, statusCode int) *BaseHttpResponse {
	return &BaseHttpResponse{
		Status:     "success",
		StatusCode: statusCode,
		Data:       data,
	}
}

func GenerateErrorResponse(err error, statusCode int, extra string) *BaseErrorResponse {
	if extra == "" {
		extra = "" // Set extra to an empty string if it's not provided
	}
	return &BaseErrorResponse{
		Status:     "error",
		StatusCode: statusCode,
		Message:    err.Error(),
		Extra:      extra,
	}
}