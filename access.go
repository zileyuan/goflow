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

func InitAccess() {
	if orm == nil {
		log.Info(DbDriverConnstr)
		connString := fmt.Sprintf(DbDriverConnstr, DbUsername, DbPassword,
			DbServer, DbPort, DbDatebase)

		log.Info(connString)
		var err error
		orm, err = xorm.NewEngine(DbDriver, connString)
		fmt.Printf(connString)

		if err != nil {
			log.Errorf("fail to init engine %v", err)
			panic(fmt.Errorf("fail to init engine!"))
		}

		orm.TZLocation = time.Local
		orm.ShowSQL = true

		tbMapper := core.NewPrefixMapper(core.SameMapper{}, "GF_")
		orm.SetTableMapper(tbMapper)
		orm.SetColumnMapper(core.SameMapper{})

		//orm.DumpAllToFile("./db_struct.sql")
		orm.Sync2(new(HistoryOrder), new(HistoryTask), new(HistoryTaskActor),
			new(Order), new(Process), new(Surrogate), new(Task), new(TaskActor))
	}
}

func Save(inf interface{}, id interface{}) error {
	session := orm.NewSession()
	defer session.Close()
	_, err := session.InsertOne(inf)
	t := reflect.TypeOf(inf)
	log.Infof(t.Name()+" %v inserted", id)
	return err
}

func Update(inf interface{}, id interface{}) error {
	session := orm.NewSession()
	defer session.Close()
	_, err := session.Id(id).Update(inf)
	t := reflect.TypeOf(inf)
	log.Infof(t.Name()+" %v updated", id)
	return err
}

func DeleteById(inf interface{}, id interface{}) error {
	session := orm.NewSession()
	defer session.Close()
	_, err := session.Id(id).Delete(inf)
	t := reflect.TypeOf(inf)
	log.Infof(t.Name()+" %v deleted", id)
	return err
}

func Delete(inf interface{}) error {
	session := orm.NewSession()
	defer session.Close()
	_, err := session.Delete(inf)
	t := reflect.TypeOf(inf)
	log.Info(t.Name() + " deleted")
	return err
}
