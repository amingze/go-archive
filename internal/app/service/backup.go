package service

import "go-archive/internal/app/dao"

type Backup struct {
	d *dao.Backup
}

func NewBackup() *Backup {
	return &Backup{dao.NewBackup()}
}
