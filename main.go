package main

import (
	"SIMS/config"
	"SIMS/global"
	orm "SIMS/init/database"
	"SIMS/init/globalID"
	initTableStruct "SIMS/init/tableStruct"
	"SIMS/init/trans"
	"SIMS/router"
)

func main() {
	orm.InitMySQL(config.AdminConfig.MysqlAdmin)
	sqlDB, _ := global.GDB.DB()
	defer sqlDB.Close()
	err := globalID.Init(1)
	if err != nil {
		panic("ID生成器初始化失败")
	}
	initTableStruct.InitTableStruct(global.GDB)
	err = trans.InitTrans("zh")
	if err != nil {
		return
	}
	_ = router.InitRouter().Run()
}
