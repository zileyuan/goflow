package goflow

import (
	"fmt"
	"reflect"
	"time"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"github.com/lunny/log"
)

var orm *xorm.Engine

//初始化数据库ORM引擎
func InitAccessByXorm(xorm *xorm.Engine) {
	orm = xorm
	//orm.DumpAllToFile("./db_struct.sql")
	err := orm.Sync2(new(HistoryOrder), new(HistoryTask), new(HistoryTaskActor),
		new(Order), new(Process), new(Surrogate), new(Task), new(TaskActor), new(CCOrder))

	PanicIf(err, "fail to sync tables")
}

//初始化数据库ORM引擎
func InitAccessByConfig(cfg string) {
	InitConfig(cfg)
	if orm == nil {
		log.Info(DbDriverConnstr)
		connString := fmt.Sprintf(DbDriverConnstr, DbUsername, DbPassword,
			DbServer, DbPort, DbDatebase)

		log.Info(connString)
		var err error
		orm, err = xorm.NewEngine(DbDriver, connString)
		fmt.Printf(connString)

		PanicIf(err, "fail to init engine")

		orm.TZLocation = time.Local
		orm.ShowSQL = true

		tbMapper := core.NewPrefixMapper(core.SameMapper{}, "GF_")
		orm.SetTableMapper(tbMapper)
		orm.SetColumnMapper(core.SameMapper{})

		InitAccessByXorm(orm)
	}
}

//保存实体对象
func Save(inf interface{}, id interface{}) {
	session := orm.NewSession()
	defer session.Close()
	_, err := session.InsertOne(inf)
	t := reflect.TypeOf(inf)
	PanicIf(err, "fail to insert %v %v", t, id)
	log.Infof("%v %v inserted", t, id)
}

//更新实体对象
func Update(inf interface{}, id interface{}) {
	session := orm.NewSession()
	defer session.Close()
	_, err := session.Id(id).Update(inf)
	t := reflect.TypeOf(inf)
	PanicIf(err, "fail to update %v %v", t, id)
	log.Infof("%v %v updated", t, id)
}

//删除实体对象
func Delete(inf interface{}, id interface{}) {
	session := orm.NewSession()
	defer session.Close()
	_, err := session.Id(id).Delete(inf)
	t := reflect.TypeOf(inf)
	PanicIf(err, "fail to delete %v %v", t, id)
	log.Infof("%v %v deleted", t, id)
}

//删除实体对象
func DeleteObj(inf interface{}) {
	session := orm.NewSession()
	defer session.Close()
	_, err := session.Delete(inf)
	t := reflect.TypeOf(inf)
	PanicIf(err, "fail to delete %v", t)
	log.Info("%v deleted", t)
}
