package utils

type ApiResponse struct {
	Status int            `json:"status"`
	Data   map[string]any `json:"data"`
}

func ResponseMapper(status int, data *map[string]any) ApiResponse {
	return ApiResponse{Status: status, Data: *data}
}
