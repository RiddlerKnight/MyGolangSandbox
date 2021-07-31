package repositories

import (
	diinterfaces "GodSpeed/Simple/di/di_interfaces"
	"fmt"
)

type simpleRepository struct {
}

func NewSimpleRepository() (diinterfaces.ISimpleRepository, error) {
	simpleRepo := simpleRepository{}
	return &simpleRepo, nil
}

func (repo simpleRepository) Find(entity interface{}) {
	fmt.Println("finding", entity)
}

func (repo simpleRepository) Update(entity interface{}) {
	fmt.Println("updating", entity)
}

func (repo simpleRepository) Del(entity interface{}) {
	fmt.Println("del", entity)
}
