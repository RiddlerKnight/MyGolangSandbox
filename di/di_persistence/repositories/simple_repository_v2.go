package repositories

import "fmt"

type simpleRepositoryV2 struct {
}

func NewSimpleRepositoryV2() simpleRepositoryV2 {
	return simpleRepositoryV2{}
}

func (repo simpleRepositoryV2) Find(entity interface{}) {
	fmt.Println("finding v2", entity)
}

func (repo simpleRepositoryV2) Update(entity interface{}) {
	fmt.Println("updating v2", entity)
}

func (repo simpleRepositoryV2) Del(entity interface{}) {
	fmt.Println("del v2", entity)
}
