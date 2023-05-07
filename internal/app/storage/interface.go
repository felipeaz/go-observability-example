package storage

type Storage interface {
	Set(id string, value interface{}) error
	Get(id string) ([]byte, error)
	Remove(id string) error
}
