package middlewares

import (
	"blogapi/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticateJWT() gin.HandlerFunc{
	return func (context *gin.Context)  {
		authHeader:=context.GetHeader("Authorization")
		if authHeader==""{
			context.JSON(http.StatusUnauthorized,utils.GetErrorResponse("Token is required"))
			context.Abort()
		}
		bearerToken:=strings.Split(authHeader," ")
		if(len(bearerToken)!=2||strings.ToLower(bearerToken[0])!="bearer"){
			context.JSON(http.StatusUnauthorized,utils.GetErrorResponse("Bearer token is required"))
			context.Abort()
		}
		tokenString:=bearerToken[1]
		userId,err:=utils.GetUserIdByToken(tokenString)
		if err!=nil{
			context.JSON(http.StatusUnauthorized, utils.GetErrorResponse("Token invalid"))
			context.Abort()
		}
		context.Set("userId",userId)
		context.Next()
	}
}