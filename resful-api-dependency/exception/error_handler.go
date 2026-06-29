package exception

import (
	"golang-restful-api/helper"
	"golang-restful-api/model/web"
	"net/http"

	"github.com/go-playground/validator"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}
	internalServerError(writer, request, err)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.HttpResponse{
			Code:   http.StatusNotFound,
			Status: "NOT_FOUND",
			Data:   exception.Error,
		}
		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.HttpResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   exception.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.HttpResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
