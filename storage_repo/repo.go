package storage_repo

import "context"

var storage_repo StorageRepo = nil

func GetStorageRepo(ctx context.Context) StorageRepo {
	if storage_repo == nil {
		storage_repo = &StorageRepoSqlite3{}
		storage_repo.Init(ctx)
	}
	return storage_repo
}
