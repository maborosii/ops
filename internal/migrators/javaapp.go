package migrators

import (
	"fmt"
	"io"
	"ops/internal/apps"
	"os"
	"path"
)

type JavaMigrator struct {
	*BaseMigrator
}

// backup old package
func (j *JavaMigrator) OldPackageMovement(a apps.AppInfo) error {
	packagePath := path.Join(a.GetBuildPath(), a.GetExtraPath(), a.GetName())
	backupPath := path.Join(j.GetBackupPath(), a.GetName())
	if !j.IsExistsBackupPath() {
		return fmt.Errorf("backup path is not exists")
	}
	err := os.Rename(packagePath, backupPath)
	return err
}

// migrate new package
func (j *JavaMigrator) NewPackageMovement(a apps.AppInfo) error {
	packagePath := path.Join(a.GetBuildPath(), a.GetExtraPath(), a.GetName())
	newPath := path.Join(j.GetNewPath(), a.GetName())
	if !j.IsExistsNewPath() {
		return fmt.Errorf("upload path is not exists")
	}

	var err error
	srcFile, err := os.Open(newPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(packagePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

// exec migrator
func (j *JavaMigrator) Run(a apps.AppInfo) {
	go func() {
		if !j.IsExistsBackupPath() {
			j.CreateBackupPath()
		}
		err := j.OldPackageMovement(a)
		if err != nil {
			return
		}
		err = j.NewPackageMovement(a)
		if err != nil {
			return
		}
		j.outChan <- a
	}()

}
