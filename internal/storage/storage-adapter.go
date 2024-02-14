package storage

type StorageAdapter interface {
	CountRequests(svcName string) int64
	StoreRequest(svcName string) (string, error)
}
