package db

import (
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/define"
	"github.com/esonhugh/tencent-coding-openapi/utils/Error"
	"github.com/glebarez/sqlite"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
)

var GlobalDatabase *GlobalDB

type GlobalDB struct {
	MainDB *gorm.DB
}

func Open(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		_, err = os.Create(path)
	}
	Error.HandleError(err)
	return db
}

func Init(DB string) {
	GlobalDatabase = new(GlobalDB)
	GlobalDatabase.MainDB = Open(DB)
	Error.HandleError(GlobalDatabase.MainDB.AutoMigrate(&define.ProjectObject{}))
	Error.HandleError(GlobalDatabase.MainDB.AutoMigrate(&define.DeoptObject{}))
	Error.HandleError(GlobalDatabase.MainDB.AutoMigrate(&define.MemberObject{}))
	log.Println("Database Init Successfully")
}
