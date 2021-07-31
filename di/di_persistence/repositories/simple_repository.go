package repositories

import (
	"fmt"
)

type simpleRepository struct {
}

func NewSimpleRepository() simpleRepository {
	return simpleRepository{}
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
