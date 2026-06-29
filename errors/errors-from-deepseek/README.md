Professional Go programmers handle errors in RESTful APIs by ensuring consistent error responses, appropriate HTTP status codes, and structured JSON error formats. Here's a comprehensive approach with examples:

### 1. **Define a Standard Error Response Structure**
All error responses follow a consistent JSON structure:

```go
type ErrorResponse struct {
    Error struct {
        Code    string      `json:"code"`
        Message string      `json:"message"`
        Details interface{} `json:"details,omitempty"`
    } `json:"error"`
}
```

### 2. **Custom Error Type with Metadata**
Create a custom error type to include HTTP status codes and error details:

```go
type APIError struct {
    StatusCode int
    Code       string
    Message    string
    Details    interface{}
}

func (e *APIError) Error() string {
    return e.Message
}
```

### 3. **Middleware for Error Handling**
Middleware captures errors, recovers from panics, and formats responses:

```go
func ErrorMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if rec := recover(); rec != nil {
                log.Printf("panic: %v", rec)
                WriteErrorResponse(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Internal server error", nil)
            }
        }()

        next.ServeHTTP(w, r)
    })
}
```

### 4. **Helper Function to Write Errors**
A helper standardizes error responses:

```go
func WriteErrorResponse(w http.ResponseWriter, statusCode int, code, message string, details interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(ErrorResponse{
        Error: struct {
            Code    string      `json:"code"`
            Message string      `json:"message"`
            Details interface{} `json:"details,omitempty"`
        }{
            Code:    code,
            Message: message,
            Details: details,
        },
    })
}
```

### 5. **Handler Examples with Error Scenarios**

#### **Example 1: Resource Not Found (404)**
```go
func getUserHandler(w http.ResponseWriter, r *http.Request) {
    userID := mux.Vars(r)["id"]
    user, err := database.GetUser(userID)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            WriteErrorResponse(w, http.StatusNotFound, "USER_NOT_FOUND", "User not found", nil)
            return
        }
        WriteErrorResponse(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Database error", nil)
        return
    }
    json.NewEncoder(w).Encode(user)
}
```

**JSON Response:**
```json
{
    "error": {
        "code": "USER_NOT_FOUND",
        "message": "User not found"
    }
}
```

#### **Example 2: Validation Error (400)**
```go
func createUserHandler(w http.ResponseWriter, r *http.Request) {
    var req UserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        WriteErrorResponse(w, http.StatusBadRequest, "INVALID_REQUEST", "Invalid JSON body", nil)
        return
    }

    validate := validator.New()
    if err := validate.Struct(req); err != nil {
        details := []map[string]string{}
        for _, err := range err.(validator.ValidationErrors) {
            details = append(details, map[string]string{
                "field":   err.Field(),
                "message": fmt.Sprintf("Field validation failed: %s", err.Tag()),
            })
        }
        WriteErrorResponse(w, http.StatusBadRequest, "VALIDATION_ERROR", "Validation failed", details)
        return
    }
    // Proceed to create user...
}
```

**JSON Response:**
```json
{
    "error": {
        "code": "VALIDATION_ERROR",
        "message": "Validation failed",
        "details": [
            {
                "field": "Email",
                "message": "Field validation failed: email"
            },
            {
                "field": "Password",
                "message": "Field validation failed: min"
            }
        ]
    }
}
```

#### **Example 3: Internal Server Error (500)**
Triggered by unexpected issues (e.g., database connection failure).

**JSON Response:**
```json
{
    "error": {
        "code": "INTERNAL_ERROR",
        "message": "Internal server error"
    }
}
```

### 6. **Panic Recovery**
The middleware recovers from panics and returns a `500` error:

```go
defer func() {
    if rec := recover(); rec != nil {
        log.Printf("panic: %v", rec)
        WriteErrorResponse(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Internal server error", nil)
    }
}()
```

### 7. **Logging for Debugging**
Log errors with context (e.g., request ID, timestamp) without exposing sensitive data:

```go
log.Printf("error: %s (code: %s, status: %d)", err.Message, err.Code, err.StatusCode)
```

### Key Takeaways:
- **Consistency**: All errors use the same JSON structure.
- **HTTP Status Codes**: Correct status codes (e.g., `404` for missing resources).
- **Error Types**: Custom errors carry metadata (e.g., status code, user-friendly message).
- **Validation Details**: Return specific field-level errors for bad requests.
- **Security**: Avoid exposing stack traces or sensitive data in production.

This approach ensures robust, user-friendly, and maintainable error handling in Go REST APIs.

source [deepseek](https://chat.deepseek.com/a/chat/s/c0c14a14-33b6-492d-9e56-5dfb7bce45f5)
