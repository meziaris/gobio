package helper

import "gobio/app/model"

func APIResponse(message string, code int, status string, data interface{}) model.ApiResponse {
	meta := model.MetaResponse{
		Message: message,
		Code:    code,
		Status:  status,
	}
	response := model.ApiResponse{
		Meta: meta,
		Data: data,
	}
	return response
}
