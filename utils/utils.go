package utils

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	ReturnCode string      `json:"returnCode"`
	ReturnMsg  string      `json:"returnMsg"`
	Data       interface{} `json:"data"`
}

func Fail(context echo.Context, returnCode string, returnMsg string) error {
	var res = Response{
		ReturnCode: returnCode,
		ReturnMsg:  returnMsg,
	}
	return context.JSON(http.StatusOK, res)
}

func Ok(context echo.Context, data interface{}) error {
	var res = Response{
		ReturnCode: "0000",
		ReturnMsg:"Success!",
		Data:       data,
	}
	return context.JSON(http.StatusOK, res)
}

// 字符串截取函数
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}