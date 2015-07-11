package goflow

import (
	"fmt"

	"github.com/Unknwon/goconfig"
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
	MaxIdleConns    int
	MaxOpenConns    int
)

//系统配置文件初始化（包含运行模式和数据库）
func InitConfig(cfg string) {
	if Cfg == nil {
		var err error
		Cfg, err = goconfig.LoadConfigFile(cfg)
		if err != nil {
			panic(fmt.Errorf("fail to load config file '%s': %v", cfg, err))
		}

		RunMode = Cfg.MustValue("app", "run_mode", "dev")
		DbDriver = Cfg.MustValue(RunMode, "db_driver")
		DbDriverConnstr = Cfg.MustValue(RunMode, "db_driver_connstr")
		DbUsername = Cfg.MustValue(RunMode, "db_username")
		DbPassword = Cfg.MustValue(RunMode, "db_password")
		DbServer = Cfg.MustValue(RunMode, "db_server")
		DbDatebase = Cfg.MustValue(RunMode, "db_datebase")
		DbPort = Cfg.MustInt(RunMode, "db_port")
		MaxIdleConns = Cfg.MustInt(RunMode, "max_idle_conns")
		MaxOpenConns = Cfg.MustInt(RunMode, "max_open_conns")
	}
}
