package dao

import (
	"go-archive/internal/app/model"
	"time"
)

type Backup struct{}

func NewBackup() *Backup {
	return &Backup{}
}

func (Backup) Create(backup model.Backup) (id int, err error) {
	id, err = gdb.Insert(&backup)
	return id, err
}

func (Backup) Find(id int) (backup model.Backup, err error) {
	err = gdb.Find(id, &backup)
	return
}

func (Backup) Delete(id int) (err error) {
	err = gdb.Delete(model.BackupTableName, id)
	return
}

func (d Backup) UpdateStatus(id int, status model.BackupStatus) (result model.Backup, err error) {
	err = gdb.Find(id, &result)
	if err != nil {
		return result, err
	}
	result.Status = status
	err = gdb.Update(&result)
	return result, err
}

func (d Backup) ChangeStatusFinish(id int) (result model.Backup, err error) {
	err = gdb.Find(id, &result)
	if err != nil {
		return result, err
	}
	result.Status = model.BackupFinish
	result.FinishAt = time.Now()
	err = gdb.Update(&result)
	return result, err
}
