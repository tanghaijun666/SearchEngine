package commom

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB

func init() {
	var err error
	// dsn := "root:710382941@tcp(127.0.0.1:3306)/tiktok?charset=utf8&parseTime=True&loc=Local&timeout=10s"
	//dsn := "root:Nulixuexi123!@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local&timeout=10s"
	//dsn := "root:wodeakun123@tcp(127.0.0.1:3306)/b_project-qingxunying-mytiktok?charset=utf8&parseTime=True&loc=Local&timeout=10s"
	dsn := "root:wodemysql123...@tcp(127.0.0.1:3306)/b_project-qingxunying-mytiktok?charset=utf8&parseTime=True&loc=Local&timeout=10s"
	_db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("数据库连接错误,error=" + err.Error())
	}
	sqlDB, _ := _db.DB()
	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}

func GetDB() *gorm.DB {
	return _db
}
