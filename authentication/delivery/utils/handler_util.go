package utils

import (
	"MPPLProject/authentication/utils"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)

	switch err {
	case utils.ErrInternalServerError:
		return http.StatusInternalServerError
	case utils.ErrNotFound:
		return http.StatusNotFound
	case utils.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}