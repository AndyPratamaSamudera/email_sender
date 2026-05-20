package models

type SendEmailRequest struct {
	AccessCode     string `form:"accessCode" json:"accessCode" binding:"required" example:"your_access_code_here"`
	EmailSender    string `form:"emailSender" json:"emailSender" binding:"required,email" example:"your-email@example.com"`
	SenderPassword string `form:"senderPassword" json:"senderPassword" binding:"required" example:"your_app_password"`
	EmailRecipient string `form:"emailRecipient" json:"emailRecipient" binding:"required,email" example:"recipient@example.com"`
	Subject        string `form:"subject" json:"subject" binding:"required" example:"Testing Email"`
	Body           string `form:"body" json:"body" binding:"required" example:"<html><body><h1>Hello World</h1></body></html>"`
}

type SuccessResponse struct {
	IsSuccess bool   `json:"is_success"`
	Message   string `json:"message"`
}

type ErrorResponse struct {
	IsSuccess bool   `json:"is_success"`
	Message   string `json:"message"`
}

type BaseResponse struct {
	Message string `json:"message"`
}
