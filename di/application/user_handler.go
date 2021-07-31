package application

import (
	diinterfaces "GodSpeed/Simple/di/di_interfaces"
)

type UserHandler struct {
	Repo diinterfaces.IBaseRepository `name: "SimpleRepoV1"`
}

func NewUserHandlerV1(repo diinterfaces.IBaseRepository) UserHandler {
	return UserHandler{Repo: repo}
}
