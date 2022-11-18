package storage_repo

import (
	"context"
	"telegraph/config"
)

var storage_repo StorageRepo = nil

func GetStorageRepo(ctx context.Context) StorageRepo {
	if storage_repo == nil {
		switch config.GetConfigs().StorageType {
		case "mongo":
			storage_repo = &StorageRepoMongo{}
		case "sqlite3":
			storage_repo = &StorageRepoSqlite3{}
		default: // default sqlite3
			storage_repo = &StorageRepoSqlite3{}
		}
		storage_repo.Init(ctx)
	}
	return storage_repo
}
