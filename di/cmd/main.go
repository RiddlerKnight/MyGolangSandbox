package main

import (
	"GodSpeed/Simple/di/application"
	diinterfaces "GodSpeed/Simple/di/di_interfaces"
	"GodSpeed/Simple/di/di_persistence/repositories"
	"GodSpeed/Simple/di/dto"
)

func main() {
	// This is normaly use interface
	var simple_repoV1 diinterfaces.IBaseRepository = repositories.NewSimpleRepository()
	userDto := dto.User{ID: 1, FirstName: "John", LastName: "Wick", Age: 32}
	application.InsertUser(simple_repoV1, &userDto)

	var simple_repoV2 diinterfaces.IBaseRepository = repositories.NewSimpleRepositoryV2()
	userDto2 := dto.User{ID: 1, FirstName: "John", LastName: "Wick", Age: 32}
	application.InsertUser(simple_repoV2, &userDto2)
}
