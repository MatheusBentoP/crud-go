package service

import (
	"fmt"

	"github.com/MatheusBentoP/crud-go/src/configuration/logger"
	"github.com/MatheusBentoP/crud-go/src/configuration/rest_err"
	"github.com/MatheusBentoP/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterfae,
) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}
