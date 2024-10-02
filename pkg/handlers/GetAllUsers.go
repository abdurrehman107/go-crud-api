package handlers

import (
	"github.com/gin-gonic/gin"
	mocks "go-rest-api/pkg/mocks"
	"net/http"
)

func GetUsers(ctx *gin.Context){
	ctx.JSON(http.StatusOK, mocks.Users)
}
