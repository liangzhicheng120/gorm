package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm/model"
)

func main() {

	// 保存所有字段,Save 会保存所有的字段,即使字段是零值
	dbClient := CreateDbClient()
	user := model.UserTab{Name: "lzc", Age: 100}
	result := dbClient.Where("id = ?", 1).Save(&user)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)

	// 更新单个列, 使用Update，需要指定条件，否则会返回 ErrMissingWhereClause 错误
	dbClient.Model(&user).Update("name", "lzc1")

	// Updates 方法支持 struct 和 map[string]interface{} 参数。
	// 当使用 struct 更新时，默认情况下，GORM 只会更新非零值的字段
	// 注意 当通过 struct 更新时，GORM 只会更新非零字段。
	// 如果想确保指定字段被更新，应该使用 Select 更新选定字段，或使用 map 来完成更新操作
	dbClient.Model(&user).Where("id = ?", 1).Updates(&user)

	// 如果想要在更新时选定、忽略某些字段，可以使用 Select、Omit
	dbClient.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18})

	dbClient.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18})

	// 批量更新 , 根据 struct 更新, 根据 map 更新
	dbClient.Model(model.UserTab{}).Where("id = ?", 1).Updates(&user)

	dbClient.Table("users").Where("id IN ?", []int{1, 2}).Updates(map[string]interface{}{"name": "lzc1", "age": 18})

	// 如果在没有任何条件的情况下执行批量更新，默认情况下，GORM 不会执行该操作，并返回 ErrMissingWhereClause 错误
	// 可以启用 AllowGlobalUpdate 模式进行全局更新或者使用原生SQL
	dbClient.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&model.UserTab{}).Update("name", "lzc")

	// 允许使用 SQL 表达式更新列
	dbClient.Model(&model.UserTab{}).Where("id = ?", 1).UpdateColumn("member_number", gorm.Expr("member_number + ?", 1))

	// 允许使用子查询更新表
	dbClient.Model(&model.UserTab{}).Update("email", dbClient.Table("user_tab").Select("email").Where("id = ?", 1))

}
