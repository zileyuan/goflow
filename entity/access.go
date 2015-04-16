package entity

import (
	"fmt"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"github.com/lunny/log"
	"goflow/cfg"
	"time"
)

var orm *xorm.Engine

func init() {
	if orm == nil {
		connString := fmt.Sprintf(cfg.DbDriverConnstr, cfg.DbServer,
			cfg.DbPort, cfg.DbUsername, cfg.DbPassword, cfg.DbDatebase)

		log.Info(connString)
		var err error
		orm, err = xorm.NewEngine(cfg.DbDriver, connString)

		if err == nil {
			orm.TZLocation = time.Local
			orm.ShowSQL = true
			orm.SetMapper(core.SameMapper{})
		} else {
			log.Error("%v", err)
		}
	}
	if orm == nil {
		panic(fmt.Errorf("fail to init engine!"))
	} else {
		//defaultEngine.DumpAllToFile("./db_struct.sql")
		orm.Sync2(new(HistoryOrder), new(HistoryTask), new(HistoryTaskActor),
			new(Order), new(Process), new(Surrogate), new(Task), new(TaskActor))
	}
}
