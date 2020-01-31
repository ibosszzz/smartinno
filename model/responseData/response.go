package responseData

type ResponseResult struct {
	Status int `json:"status"`
	Error  string `json:"error"`
	Data interface{} `json:"data"`
}

type ResponseResultExcel struct {
	Status int `json:"status"`
	Error  []error `json:"error"`
	Data interface{} `json:"data"`
}

type ResponseResultExcelWithError struct {
	Status int `json:"status"`
	Error  []string `json:"error"`
	Data interface{} `json:"data"`
}