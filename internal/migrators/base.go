package migrators

import (
	"ops/internal/apps"
	"os"
)

type MigrateItf interface {
	IsExistsBackupPath() bool
	IsExistsNewPath() bool
	CreateBackupPath() error
	GetBackupPath() string
	GetNewPath() string

	OldPackageMovement(a apps.AppInfo) error
	NewPackageMovement(a apps.AppInfo) error

	SetOutChan(chan<- apps.AppInfo)
	Run(apps.AppInfo)
}

type BaseMigrator struct {
	outChan    chan<- apps.AppInfo
	backupPath string
	newPath    string
}

func (b *BaseMigrator) GetNewPath() string {
	return b.newPath
}

func (b *BaseMigrator) GetBackupPath() string {
	return b.backupPath
}

func (b *BaseMigrator) IsExistsBackupPath() bool {
	_, err := os.Stat(b.backupPath)
	if err != nil {
		return false
	}
	return true
}

func (b *BaseMigrator) IsExistsNewPath() bool {
	_, err := os.Stat(b.newPath)
	if err != nil {
		return false
	}
	return true
}

func (b *BaseMigrator) CreateBackupPath() error {
	return os.Mkdir(b.backupPath, 0644)
}

func (b *BaseMigrator) SetOutChan(outchan chan<- apps.AppInfo) {
	b.outChan = outchan
}
