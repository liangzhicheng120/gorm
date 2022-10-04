package main

import (
	"gorm/model"
)

func main4() {

	dbClient := CreateDbClient()

	// 更新单个列, 使用Update，需要指定条件，否则会返回 ErrMissingWhereClause 错误
	//dbClient.Model(&model.UserTab{}).Where("id = ?",1).Update("name", "1")

	// Updates 方法支持 struct 和 map[string]interface{} 参数。
	// 当使用 struct 更新时，默认情况下，GORM 只会更新非零值的字段
	// 注意 当通过 struct 更新时，GORM 只会更新非零字段。
	// 如果想确保指定字段被更新，应该使用 Select 更新选定字段，或使用 map 来完成更新操作
	//user := model.UserTab{Name: "1"}
	//dbClient.Model(&model.UserTab{}).Where("id = ?", 1).Updates(map[string]interface{}{
	//	"name": "123",
	//})
	//
	//// 如果想要在更新时选定、忽略某些字段，可以使用 Select、Omit
	//dbClient.Model(&model.UserTab{}).Where("id = ?", 1).
	//	Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18})

	//dbClient.Model(&model.UserTab{}).Omit("name").Where("id = ?", 1).
	//	Updates(map[string]interface{}{"name": "hello", "age": 10})

	//// 批量更新 , 根据 struct 更新, 根据 map 更新
	//dbClient.Model(model.UserTab{}).Where("id = ?", 1).Updates(&user)
	//
	//dbClient.Table("user_tab").Where("id in ?", []int{1, 2}).Updates(map[string]interface{}{"name": "lzc1", "age": 20})
	//
	//// 如果在没有任何条件的情况下执行批量更新，默认情况下，GORM 不会执行该操作，并返回 ErrMissingWhereClause 错误
	//// 可以启用 AllowGlobalUpdate 模式进行全局更新或者使用原生SQL
	//dbClient.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&model.UserTab{}).Update("name", "lzc")
	//
	//// 允许使用 SQL 表达式更新列
	//dbClient.Model(&model.UserTab{}).Where("id = ?", 1).UpdateColumn("member_number", gorm.Expr("?", 3))
	//
	//// 允许使用子查询更新表
	dbClient.Model(&model.UserTab{}).Where("id = ?", 1).Update("email", dbClient.Table("user1_tab").Select("email").Where("id = ?", 1))

}
