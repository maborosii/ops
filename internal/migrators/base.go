package migrators

type MigrateItf interface {
	IsExistsBackupPath(path string) bool
	CreateBackupPath(path string) error
	GetBackupPath() string
	GetNewPath(path string) string

	OldPackageMovement() error
	NewPackageMovement() error
}
