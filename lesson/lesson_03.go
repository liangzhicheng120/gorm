package main

import (
	"gorm.io/gorm"
	"gorm/model"
)

func main3() {
	dbClient := CreateDbClient()

	// GORM 提供了 First、Take、Last 方法，从数据库中检索单个对象。
	// 当查询数据库时它添加了 LIMIT 1 条件，且没有找到记录时，它会返回 ErrRecordNotFound 错误
	// First 和 Last 会根据主键排序，分别查询第一条和最后一条记录。
	// 如果 model 没有定义主键，那么将按 model 的第一个字段进行排序

	// 获取第一条记录（主键升序）
	// SELECT * FROM user_tab ORDER BY id LIMIT 1;
	dbClient.First(&model.UserTab{})

	// 获取一条记录，没有指定排序字段
	// SELECT * FROM user_tab LIMIT 1;
	dbClient.Take(&model.UserTab{})

	// 获取最后一条记录（主键降序）
	// SELECT * FROM user_tab ORDER BY id DESC LIMIT 1;
	dbClient.Last(&model.UserTab{})

	// 如果想避免ErrRecordNotFound错误，可以使用Find
	dbClient.Limit(1).Find(&model.UserTab{})

	// 条件查询
	var result model.UserTab
	dbClient.Where("name = ? and is_del = ?", "lzc", 0).Take(&result)

	// 根据结构体查询
	// 当使用struct查询时，GORM只对非零字段进行查询，字段的值是0，''，false或其他零值，它将不会被用来建立查询条件
	dbClient.Where(&model.UserTab{Name: "lzc", Age: 10}).Find(&result)

	// 如果想要包含零值查询条件，可以使用 map，其会包含所有 key-value 的查询条件
	dbClient.Where(map[string]interface{}{"Name": "lzc", "Age": 10}).Find(&result)

	// 内联条件查询
	dbClient.Find(&result, "name = ? and is_del = ?", "lzc", 0)

	// not 查询
	var users []model.UserTab
	dbClient.Not("name = ?", "lzc").Find(&users)

	// not in 查询
	dbClient.Not(map[string]interface{}{"name": "lzc"}).Find(&users)

	// or 查询
	dbClient.Where("name = ?", "lzc").Or("name = ?", "lzc1").Find(&users)

	// Select 允许指定从数据库中检索哪些字段，默认情况下，GORM 会检索所有字段
	dbClient.Select("name", "age").Find(&users)

	// 排序查询
	dbClient.Order("age desc, name").Find(&users)

	// 分页查询，也可以自定义分页器
	dbClient.Offset(0).Limit(10).Find(&users)
	dbClient.Scopes(Paginate(1, 10)).Find(&users)

	// Group By & Having 查询
	dbClient.Model(&model.UserTab{}).Group("name").Find(&users)

	dbClient.Model(&model.UserTab{}).Group("name").
		Having("name = ?", "lzc").Find(&result)

	// Distinct 查询
	dbClient.Distinct("name", "age").Find(&users)

	// join 查询
	dbClient.Table("user_tab as u").
		Joins("left join user1_tab u on u1.id = u.id").Find(&users)

}

func Paginate(currentPage, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if currentPage == 0 {
			currentPage = 1
		}
		switch {
		case pageSize > 10000:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (currentPage - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
