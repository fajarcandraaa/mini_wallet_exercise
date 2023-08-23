package dto

import (
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/presentation"
)

func ToResponse(status string, data interface{}) presentation.Response {
	res := presentation.Response{
		Status: status,
		Data:   data,
	}

	return res
}