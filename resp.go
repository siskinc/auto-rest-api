package autorestapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data    interface{} `json:"data"`
	ErrMsg  string      `json:"err_msg"`
	Success int         `json:"success"`
}

type ResponseList struct {
	Data    interface{} `json:"data,omitempty" bson:"data"`
	Count   int         `json:"count,omitempty" bson:"count"`
	ErrMsg  string      `json:"err_msg,omitempty" bson:"err_msg"`
	Success int         `json:"success,omitempty" bson:"success"`
}

func RespErr(c *gin.Context, errMsg string) {
	resp := &Response{}
	resp.Success = 0
	resp.ErrMsg = errMsg
	c.JSON(http.StatusOK, resp)
}

func RespData(c *gin.Context, data interface{}) {
	resp := &Response{}
	resp.Success = 1
	resp.Data = data
	c.JSON(http.StatusOK, resp)
}

func RespListErr(c *gin.Context, errMsg string) {
	resp := &ResponseList{}
	resp.Success = 0
	resp.ErrMsg = errMsg
	c.JSON(http.StatusOK, resp)
}

func RespListData(c *gin.Context, data interface{}, count int) {
	resp := &ResponseList{}
	resp.Count = count
	resp.Success = 1
	resp.Data = data
	c.JSON(http.StatusOK, resp)
}
