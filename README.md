# Email Sender API

A simple API built using the **Go (Gin framework)** to send HTML emails via the SMTP protocol. This application also comes with Swagger documentation.

## Requirements
- Go (Golang) installed on your system.

## Installation & Setup
1. Make sure you are in the `email_sender` project directory.
2. Copy the environment configuration file:
   ```bash
   cp .env.example .env
   ```
   Then adjust the port value in the `.env` file (by default, the port is `8080`).
3. Download all required dependencies:
   ```bash
   go mod tidy
   ```

## Running the Application
To start the server, use the following command:
```bash
go run main.go
```
The application will start and listen for HTTP requests on the configured port.

## Swagger Documentation
Interactive documentation for this API is available and can be accessed in your browser at:
**[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)**

---

## Endpoints

### 1. Health Check
Used to verify that the API server is up and running.
- **URL**: `/`
- **Method**: `GET`
- **Success Response**:
  ```json
  {
    "message": "SMTP Email API Running"
  }
  ```

### 2. Send Email
The main endpoint for sending emails.
- **URL**: `/send-email`
- **Method**: `POST`
- **Content-Type**: `application/json` *(Also supports form-data)*

#### Example Request Body (JSON)
*(The data below is fake/mocked for demonstration purposes)*

```json
{
  "accessCode": "super-secret-x99",
  "emailSender": "notification@another-store.com",
  "senderPassword": "your-smtp-app-password",
  "emailRecipient": "budi.customer@example-email.com",
  "subject": "Year-End Special Promo!",
  "body": "<html><body><h1>Hello Budi!</h1><p>Don't miss out on our special year-end promo.</p></body></html>"
}
```

**Parameter Descriptions:**
- `accessCode`: Authorization access code required to use this endpoint.
- `emailSender`: The sender's email address (configured with SMTP/App Password).
- `senderPassword`: The application password of the sender's email.
- `emailRecipient`: The destination/recipient email address.
- `subject`: The title or subject of the email.
- `body`: The email message content, supports HTML tags.

#### Example Success Response (200 OK)
```json
{
  "is_success": true,
  "message": "Email sent successfully"
}
```

#### Example Error Response (400 Bad Request / 500 Internal Server Error)
```json
{
  "is_success": false,
  "message": "Specific error message will appear here"
}
```
