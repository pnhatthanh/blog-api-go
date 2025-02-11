package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserIdFromContext(context *gin.Context) string {
	userId, ok := context.Get("userId")
	if !ok {
		return ""
	}
	userIdStr, valid := userId.(string)
	if !valid {
		return ""
	}
	return userIdStr
}

func GetQueryInt(context *gin.Context, param string, defaultVal int) int {
	res := context.Query(param)
	val, err:=strconv.Atoi(res)
	if err!=nil{
		return defaultVal
	}
	return val
}
