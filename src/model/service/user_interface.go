package service

import (
	"github.com/MatheusBentoP/crud-go/src/configuration/rest_err"
	"github.com/MatheusBentoP/crud-go/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterfae) *rest_err.RestErr
	UptadeUser(string, model.UserDomainInterfae) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterfae, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
