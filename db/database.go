package db

import (
	"sync"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"github.com/coraldane/ops-meta/g"
)

var (
	ormLock = new(sync.RWMutex)
)

func Init() {
	// set default database
	dbConfig := g.Config().Database

	if g.Config().LogLevel == "debug" {
		orm.Debug = true
	}

	orm.RegisterDataBase("default", "mysql", dbConfig.Dsn, dbConfig.MaxIdle, dbConfig.MaxIdle)
	orm.RunSyncdb("default", false, true)
}

func NewOrm() orm.Ormer {
	ormLock.RLock()
	defer ormLock.RUnlock()
	return orm.NewOrm()
}
