package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Responder struct {
	C *gin.Context
}

type BaseResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta_data,omitempty"`
	Token   string      `json:"token,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type Pagination struct {
	TotalData int    `json:"total_data"`
	TotalPage int    `json:"total_page"`
	Page      int    `json:"page"`
	PrevLink  string `json:"prev_link,omitempty"`
	NextLink  string `json:"next_link,omitempty"`
}

type ResponseWithMeta struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    Pagination  `json:"meta_data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type LoginResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Token   string      `json:"token,omitempty"`
}

func NewResponse(ctx *gin.Context) *Responder {
	return &Responder{C: ctx}
}

func (r *Responder) Success(message string, data interface{}) {
	r.C.JSON(http.StatusOK, BaseResponse{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	})
}

func (r *Responder) SuccessWithMeta(message string, data interface{}, meta Pagination) {
	r.C.JSON(http.StatusOK, BaseResponse{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func (r *Responder) LoginSuccess(message string, data interface{}, token string) {
	r.C.JSON(http.StatusOK, LoginResponse{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
		Token:   token,
	})
}

func (r *Responder) Created(message string, data interface{}) {
	r.C.JSON(http.StatusCreated, BaseResponse{
		Status:  http.StatusCreated,
		Message: message,
		Data:    data,
	})
}

func (r *Responder) BadRequest(message string, err interface{}) {
	r.C.JSON(http.StatusBadRequest, BaseResponse{
		Status:  http.StatusBadRequest,
		Message: message,
		Error:   err,
	})
	r.C.Abort()
}

func (r *Responder) Unauthorized(message string, err interface{}) {
	r.C.JSON(http.StatusUnauthorized, BaseResponse{
		Status:  http.StatusUnauthorized,
		Message: message,
		Error:   err,
	})
	r.C.Abort()
}

func (r *Responder) Forbidden(message string, err interface{}) {
	r.C.JSON(http.StatusForbidden, BaseResponse{
		Status:  http.StatusForbidden,
		Message: message,
		Error:   err,
	})
	r.C.Abort()
}

func (r *Responder) NotFound(message string, err interface{}) {
	r.C.JSON(http.StatusNotFound, BaseResponse{
		Status:  http.StatusNotFound,
		Message: message,
		Error:   err,
	})
	r.C.Abort()
}

func (r *Responder) InternalServerError(message string, err interface{}) {
	r.C.JSON(http.StatusInternalServerError, BaseResponse{
		Status:  http.StatusInternalServerError,
		Message: message,
		Error:   err,
	})
	r.C.Abort()
}
