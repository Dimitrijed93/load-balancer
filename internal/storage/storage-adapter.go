package storage

type StorageAdapter interface {
	CountRequests(svcName string) int
	StoreRequest(svcName string, exp int) (string, error)
}
