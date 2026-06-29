package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// 1. Define a standard Error response Structure
type ErrorResponse struct {
	Error struct {
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Details interface{} `json:"details,omitempty"`
	} `json:"error"`
}

//2. Custome error Type with metadata

type APIError struct {
	StatusCode int
	Code       string
	Message    string
	Details    interface{}
}

func (e *APIError) Error() string {
	return e.Message
}

//3. Middleware for error handling

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("panic: %v", rec)
				WriterErrorResponse(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Internal server error", nil)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// 4. Helper Function to write errors
func WriterErrorResponse(w http.ResponseWriter, statusCode int, code, message string, details interface{}) {
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

// 5. handler examples with error scenarios
// example 1: resouece Not Found
func getUserHandler(w http.ResponseWriter, r *http.Request) {
  userID := mux.Vars(r)["id"]
  user, err := database.GetUser(userID)
  if err != nil {
    if errors.Is(ee, sql.ErrNoRows) {
      WriteErrorResponse(w, http.StatusNotFound, "USER_NOT_FOUND", "User not found", nil)
      return
    }
    WriterErrorResponse(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Database error", nil)
    return
  }

  json.NewEncoder(w).Encode(user)
}

// JSON response 
// {
//     "error": {
//         "code": "USER_NOT_FOUND",
//         "message": "User not found"
//     }
// }

// example 2 : validation error(400)

func createUserHandler(w http.ResponseWriter, r *http.Request) {
  var req UserRequest
  if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    WriterErrorResponse(w, http.StatusBadRequest, "INVALID_REQUEST", "Invalid JSON body", nil)
    return
  }
// validate is from go-playground/validator 
  validate := validator.New()
  if err := validate.Struct(req); err != nil {
    details := []map[string]string{}
    for _, err := range err.(validator.ValidatorErrors) {
      details = append(details, map[string]string {
        "field": err.Field(),
        "message": fmt.Sprint("Field validation failed: %s",err.Tag())
      })
    }
    WriterErrorResponse(w, http.StatusBadRequest, "VALIDATION_ERROR", "Validation failed", details)
    return 
  }
}
//
// {
//     "error": {
//         "code": "VALIDATION_ERROR",
//         "message": "Validation failed",
//         "details": [
//             {
//                 "field": "Email",
//                 "message": "Field validation failed: email"
//             },
//             {
//                 "field": "Password",
//                 "message": "Field validation failed: min"
//             }
//         ]
//     }
// }

// example 3 : internal server error(500)

// {
//     "error": {
//         "code": "INTERNAL_ERROR",
//         "message": "Internal server error"
//     }
// }

// 6. Panic recovery
// the middleware revovers from panics and return a 500 error: 

defer func() {
  if rec := recover(); rec != nil {
    log.Printf("panic: %v", rec)
    WriterErrorResponse(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Internal server error", nil)
  }
}()

// 7. Logging for debugging 
// log errirs with context (e.g., request ID, timestamp) without exposing sensitive data: 

// log.Printf("error: %s (code: %s, status: %d)", err.Message, err.Code, err.StatusCode)


