package dao

import (
	"go-archive/internal/app/model"
	"go-archive/internal/pkg/utils/fileutil"
	"go-archive/pkg/db"
	"go-archive/pkg/db/datastores/disk"
	"log"
)

var gdb *db.Database

func init() {
	fileutil.CheckAndCreateMkDir("./data")
	ds, err := disk.New("./data", ".json")
	if err != nil {
		panic(err)
	}
	gdb, err = db.New(ds)
	if err != nil {
		panic(err)
	}
	CreateTables(model.Tables()...)
}

func CreateTables(regs ...model.Register) {
	for _, reg := range regs {
		if !gdb.TableExists(reg.Table()) {
			err := gdb.CreateTable(reg.Table())
			if err != nil {
				log.Println(err.Error())
			}
		}
	}
}
