package dto

type BaseRes struct {
	Error ErrorRes    `json:"error"`
	Data  interface{} `json:"data"`
}

type ErrorRes struct {
	ErrorCode        string `form:"errorCode"`
	ErrorDescription string `form:"errorDescription"`
	ErrorMessage     string `form:"errorMessage"`
}
