package diinterfaces

type IBaseRepository interface {
	Find(interface{})
	Update(interface{})
	Del(interface{})
}
