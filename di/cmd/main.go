package main

import (
	diinterfaces "GodSpeed/Simple/di/di_interfaces"
	"GodSpeed/Simple/di/di_persistence/repositories"
	"GodSpeed/Simple/di/dto"
	"context"
	"fmt"
	"time"

	promise "github.com/fanliao/go-promise"
	"go.uber.org/fx"
)

func main() {
	// This is normaly use interface
	// var simple_repoV1 diinterfaces.IBaseRepository = repositories.NewSimpleRepository()
	// userDto := dto.User{ID: 1, FirstName: "John", LastName: "Wick", Age: 32}
	// application.InsertUser(simple_repoV1, &userDto)

	// var simple_repoV2 diinterfaces.IBaseRepository = repositories.NewSimpleRepositoryV2()
	// userDto2 := dto.User{ID: 1, FirstName: "John", LastName: "Wick", Age: 32}
	// application.InsertUser(simple_repoV2, &userDto2)

	appDi := fx.New(
		fx.Provide(
			// fx.Annotated{
			// 	Name:   "SimpleRepoV1",
			// 	Target: repositories.NewSimpleRepository},

			repositories.NewSimpleRepository,
		),

		fx.Invoke(StartApp),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := appDi.Start(startCtx); err != nil {
		fmt.Println(err)
	}
}

func StartApp(baseRepo diinterfaces.ISimpleRepository) {
	task := func() {
		for {
			baseRepo.Find(dto.User{ID: 1, FirstName: "John", LastName: "Wick"})
			time.Sleep(time.Second * 10)
		}
	}

	future := promise.WhenAll(task)
	future.Get()

}
