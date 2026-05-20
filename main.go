package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "email_sender/docs"
	"email_sender/handlers"
	"email_sender/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title SMTP Email Sender API
// @version 1.0
// @description API untuk mengirim email HTML melalui SMTP.
// @host localhost:8080
// @BasePath /

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseBodyWriter) Write(b []byte) (int, error) {
	return w.body.Write(b)
}

func swaggerCSSInjector() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/swagger/") && strings.HasSuffix(c.Request.URL.Path, "index.html") {
			w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
			c.Writer = w
			c.Next()

			css := `<style>
				.swagger-ui .topbar { display: none !important; }
				.swagger-ui .info { display: none !important; }
				.swagger-ui table.responses-table td.response-col_description div.model-example { display: none !important; }
				.swagger-ui .responses-inner .model-example { display: none !important; }
			</style>`
			modified := strings.Replace(w.body.String(), "</head>", css+"</head>", 1)
			w.ResponseWriter.Header().Del("Content-Length")
			w.ResponseWriter.Write([]byte(modified))
		} else {
			c.Next()
		}
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	_ = godotenv.Load()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.BaseResponse{
			Message: "SMTP Email API Running",
		})
	})

	r.Use(swaggerCSSInjector())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	r.POST("/send-email", handlers.SendEmailHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Link to Swagger: http://localhost:%s/swagger/index.html\n", port)
	fmt.Printf("Listening and serving HTTP on :%s\n", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
