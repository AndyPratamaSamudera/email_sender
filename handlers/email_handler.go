package handlers

import (
	"log"
	"net/http"
	"os"

	"email-sender/models"
	"email-sender/utils"

	"github.com/gin-gonic/gin"
)

// SendEmailHandler godoc
// @Summary Send HTML emails
// @Description Sending HTML email using SMTP with access code validation
// @Tags email
// @Accept json
// @Produce json
// @Param request body models.SendEmailRequest true "Request payload"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /send-email [post]
func SendEmailHandler(c *gin.Context) {
	var req models.SendEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Invalid request body: %v\n", err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			IsSuccess: false,
			Message:   "Invalid JSON format or missing fields",
		})
		return
	}

	// Validasi access code
	validHash := os.Getenv("VALID_ACCESS_CODE_HASH")
	if validHash == "" {
		log.Println("Warning: VALID_ACCESS_CODE_HASH is not set in environment")
	}

	if !utils.ValidateAccessCode(req.AccessCode, validHash) {
		log.Printf("Unauthorized access attempt with code: %s\n", req.AccessCode)
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			IsSuccess: false,
			Message:   "Unauthorized",
		})
		return
	}

	// Auto detect SMTP host
	smtpHost := utils.AutoDetectSMTPHost(req.EmailSender)
	if smtpHost == "" {
		log.Printf("Failed to auto-detect SMTP host for email: %s\n", req.EmailSender)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			IsSuccess: false,
			Message:   "SMTP host can't be auto-detected for this domain. Use gmail, outlook, or yahoo.",
		})
		return
	}

	// Set default SMTP port to 587
	smtpPort := 587

	log.Printf("Sending email from %s to %s via %s:%d\n", req.EmailSender, req.EmailRecipient, smtpHost, smtpPort)

	// Send email
	err := utils.SendEmail(
		req.EmailSender,
		req.SenderPassword,
		req.EmailRecipient,
		req.Subject,
		req.Body,
		smtpHost,
		smtpPort,
	)

	if err != nil {
		log.Printf("Failed to send email: %v\n", err)
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			IsSuccess: false,
			Message:   "Failed to send email: " + err.Error(),
		})
		return
	}

	log.Printf("Email successfully sent to %s\n", req.EmailRecipient)
	c.JSON(http.StatusOK, models.SuccessResponse{
		IsSuccess: true,
		Message:   "Email sent successfully",
	})
}
