package model

type Users interface {
	Users() map[string]UserDefinition
	Delete(id string)
	Update([]string)
	Create(user UserDefinition)
}
