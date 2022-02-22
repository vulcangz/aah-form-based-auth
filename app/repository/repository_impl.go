package repository

import (
	"aah-form-based-auth/app/database"
	"database/sql"

	aah "aahframe.work"
)

var Repo Repository

var RepoInstance Repository

// SormRepository 基于SQLBoiler的存储库封装
type SormRepository struct {
	db *sql.DB
}

// NewSormRepository 初始化并生成存储库的实例
func NewSormRepository(db *sql.DB) (Repository, error) {
	Repo := &SormRepository{
		db: db,
	}
	return Repo, nil
}

// Before method is an interceptor for admin path.
func InitRepo(_ *aah.Event) {
	aah.App().Log().Debug("Repository instance initialized is called")
	// var err error
	RepoInstance, _ = NewSormRepository(database.SDB())
}

// Before method is an interceptor for admin path.
func R() Repository {
	return RepoInstance
}
