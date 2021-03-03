package main

import (
	"SIMS/config"
	orm "SIMS/init/database"
	"SIMS/init/globalID"
	initTableStruct "SIMS/init/tableStruct"
	"SIMS/router"
)

func main() {
	db := orm.InitMySQL(config.AdminConfig.MysqlAdmin)
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	err := globalID.Init(1)
	if err != nil {
		panic("ID生成器初始化失败")
	}
	initTableStruct.InitTableStruct(db)
	_ = router.InitRouter().Run()
}
