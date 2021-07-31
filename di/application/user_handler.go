package application

import (
	diinterfaces "GodSpeed/Simple/di/di_interfaces"
	"GodSpeed/Simple/di/dto"
)

func InsertUser(repo diinterfaces.IBaseRepository, user *dto.User) {
	repo.Find(user)
}
