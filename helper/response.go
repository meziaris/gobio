package helper

import "gobio/model"

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

func LoginResponse(m model.LoginUserResponse, token string) model.LoginUserFormatterResponse {
	response := model.LoginUserFormatterResponse{
		ID:       m.ID,
		Name:     m.Name,
		Username: m.Username,
		Email:    m.Email,
		Token:    token,
	}
	return response
}
