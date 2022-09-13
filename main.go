package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm/model"
)

/*
 *           N777777777NO
 *         N7777777777777N
 *        M777777777777777N
 *        *N877777777D77777M
 *       N M77777777ONND777M
 *       MN777777777NN  D777
 *     N7ZN777777777NN ~M7778
 *    N777777777777MMNN88777N
 *    N777777777777MNZZZ7777O
 *    DZN7777O77777777777777
 *     N7OONND7777777D77777N
 *      8*M++++?N???$77777$
 *       M7++++N+M77777777N
 *        N77O777777777777$                              M
 *          DNNM$$$$777777N                              D
 *         N*N:=N$777N7777M                             NZ
 *        77Z::::N777777777                          ODZZZ
 *       77N::::::N77777777M                         NNZZZ$
 *     $777:::::::77777777MN                        ZM8ZZZZZ
 *     777M::::::Z7777777Z77                        N++ZZZZNN
 *    7777M:::::M7777777$777M                       $++IZZZZM
 *   M777$:::::N777777*M7777M                       +++++ZZZDN
 *     NN$::::::7777$*M777777N                      N+++ZZZZNZ
 *       N::::::N:7*O:77777777                      N++++ZZZZN
 *       M::::::::::::N77777777+                   +?+++++ZZZM
 *       8::::::::::::D77777777M                    O+++++ZZ
 *        ::::::::::::M777777777N                      O+?D
 *        M:::::::::::M77777777778                     77=
 *        D=::::::::::N7777777777N                    777
 *       INN===::::::=77777777777N                  I777N
 *      ?777N========N7777777777787M               N7777
 *      77777*D======N77777777777N777N?         N777777
 *     I77777$$*N7===M$$77777777$77777777*MMZ77777777N
 *      $$$$$$$$$$*NIZN$$$$$$$$*M$$7777777777777777ON
 *       M$$$$$$$*M    M$$$$$$$*N=N$$$$7777777$$*ND
 *      O77Z$$$$$$$     M$$$$$$$*MNI==*DNNNNM=~N
 *   7 :N MNN$$$*M$      $$$777$8      8D8I
 *     NMM.:7O           777777778
 *                       7777777MN
 *                       M NO .7:
 *                       M   :   M
 *                            8
 */

// Constant matcher factory methods

func main() {
	dbClient := CreateDbClient()
	if err := dbClient.Model(&model.UserTab{}).Create(&model.UserTab{
		Name:  "lzc",
		Age:   10,
		Email: "7758258@qq.com",
	}).Error; err != nil {
		fmt.Println(err)
	}
}

func CreateDbClient() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(127.0.0.1:3306)/learn_gorm_db?" +
			"charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}
