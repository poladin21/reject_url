package interfaces

type Storeges interface {
	GetByID(ID int) error
}
