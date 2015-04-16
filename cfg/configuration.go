package cfg

import (
	"fmt"

	"github.com/Unknwon/goconfig"
)

const (
	CfgPath = "conf/app.conf"
)

var (
	Cfg             *goconfig.ConfigFile
	RunMode         string
	DbDriver        string
	DbDriverConnstr string
	DbUsername      string
	DbPassword      string
	DbServer        string
	DbPort          int
	DbDatebase      string
)

func init() {
	var err error
	Cfg, err = goconfig.LoadConfigFile(CfgPath)
	if err != nil {
		panic(fmt.Errorf("fail to load config file '%s': %v", CfgPath, err))
	}

	RunMode = Cfg.MustValue("app", "run_mode", "dev")
	DbDriver = Cfg.MustValue(RunMode, "db_driver")
	DbDriverConnstr = Cfg.MustValue(RunMode, "db_driver_connstr")
	DbUsername = Cfg.MustValue(RunMode, "db_username")
	DbPassword = Cfg.MustValue(RunMode, "db_password")
	DbServer = Cfg.MustValue(RunMode, "db_server")
	DbDatebase = Cfg.MustValue(RunMode, "db_datebase")
	DbPort = Cfg.MustInt(RunMode, "db_port")
}
