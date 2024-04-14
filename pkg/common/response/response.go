package response

type Result struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

type PageResult struct {
	List     interface{} `json:"list,omitempty"`
	Current  int         `json:"current"`
	PageSize int         `json:"pageSize"`
	Total    int         `json:"total"`
}

func OK(data interface{}) *Result {
	return &Result{
		Success: true,
		Data:    data,
		Code:    0,
		Message: "OK",
	}
}

func Error(data interface{}) *Result {
	return &Result{
		Success: true,
		Data:    data,
		Code:    1,
		Message: "Error",
	}
}

func Page(data interface{}, current, pageSize, total int) *Result {
	return &Result{
		Success: true,
		Data: &PageResult{
			List:     data,
			Current:  current,
			PageSize: pageSize,
			Total:    total,
		},
		Code:    0,
		Message: "OK",
	}
}
