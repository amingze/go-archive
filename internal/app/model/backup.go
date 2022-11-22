package model

import (
	"go-archive/pkg/db"
	"time"
)

type BackupStatus int

const (
	BackupFail BackupStatus = iota - 1
	BackupNone
	BackupNoStart
	BackupUnderway
	BackupFinish
)

type Backup struct {
	ID       int          `json:"id,omitempty"`
	CreateAt time.Time    `json:"create_at,omitempty"`
	FinishAt time.Time    `json:"finish_at,omitempty"`
	Status   BackupStatus `json:"status,omitempty"`
	FileSize int64        `json:"file_size,omitempty"`
	FilePath string       `json:"file_path,omitempty"`
	FileMD5  string       `json:"file_md5,omitempty"`
}

func (a *Backup) GetID() int {
	return a.ID
}

func (a *Backup) SetID(id int) {
	a.ID = id
}

func (a *Backup) AfterFind(db *db.Database) error {
	*a = Backup(*a)
	return nil
}

func (a *Backup) Table() string {
	return "backups"
}
