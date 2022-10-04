package main

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm/model"
)

func main() {
	dbClient := CreateDbClient()

	// 如果在没有任何条件的情况下执行批量删除，GORM 不会执行该操作，并返回 ErrMissingWhereClause 错误
	dbClient.Delete(&model.UserTab{})

	// 根据主键删除
	dbClient.Delete(&model.UserTab{}, 1)

	dbClient.Where("1 = 1").Delete(&model.UserTab{})
	// 使用原生SQL
	dbClient.Exec("delete from user_tab")
	// 启用AllowGlobalUpdate模式
	dbClient.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.UserTab{})

	// 返回被删除的数据
	// 返回所有列
	var users []model.UserTab
	dbClient.Clauses(clause.Returning{}).Where("id = ?", 1).Delete(&users)

	// 返回指定的列
	dbClient.Clauses(clause.Returning{Columns: []clause.Column{{Name: "name"}, {Name: "age"}}}).
		Where("id = ?", 1).Delete(&users)

	// 使用 Unscoped 找到被软删除的记录
	dbClient.Unscoped().Where("id = ?", 2).Find(&users)

	// 使用 Unscoped 永久删除匹配的记录
	dbClient.Unscoped().Where("id = ?", 2).Delete(&model.UserTab{})

	// 使用 0 / 1 标志位进行软删除
	dbClient.Delete(&model.UserTab{}, 1)

}
