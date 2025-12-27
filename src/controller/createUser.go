package controller

import (
	"net/http"

	"github.com/MatheusBentoP/crud-go/src/configuration/validation"
	"github.com/MatheusBentoP/crud-go/src/controller/model/request"
	"github.com/MatheusBentoP/crud-go/src/model"
	"github.com/MatheusBentoP/crud-go/src/model/service"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	userDomainInterfae model.UserDomainInterfae
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zapcore.Field{
		Key:    "journey",
		String: "createUser",
	})
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to valdate user info", err)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)
	service := service.NewUserDomainService()
	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created succesfully", zap.String("journey", "createUser"))

	c.String(http.StatusOK, "")
}
