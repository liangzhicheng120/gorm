package main

import (
	"fmt"
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

func main2() {
	// INSERT INTO `user_tab` (`name`,`age`,`email`,`is_del`,`create_time`,`update_time`,`birthday`,`member_number`)
	// VALUES ('lzc',0,'',0,0,0,NULL,NULL)
	dbClient := CreateDbClient()
	user := model.UserTab{ID: 1, Name: "lzc12", Age: 15, Email: "1234567@qq.com"}
	result := dbClient.Create(&user)
	err := result.Error
	rowsAffected := result.RowsAffected
	id := user.ID
	fmt.Println(err)
	fmt.Println(rowsAffected)
	fmt.Println(id)

	// ????????????????????????????????????
	//dbClient.Select("Name", "Age").Create(&user)

	// ??????????????????,??????name , age
	//dbClient.Omit("Name", "Age").Create(&user)

	// ????????????
	//var users = []model.UserTab{{Name: "lzc1"}, {Name: "lzc2"}, {Name: "lzc3"}}
	//dbClient.Create(&users)

	// ?????? CreateInBatches ?????????????????????????????????????????????
	// ????????? 2
	//dbClient.CreateInBatches(users, 2)

	// ??????????????????
	//dbClient.Session(&gorm.Session{SkipHooks: true}).Create(&user)

	// ??????Map??????
	//dbClient.Model(&model.UserTab{}).Create(map[string]interface{}{
	//	"Name": "lzc",
	//	"Age": 18,
	//})

	// ??????????????????????????????
	//dbClient.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)

	// ????????????????????????????????????
	//dbClient.Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "id"}},
	//	DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
	//}).Create(&user)

	// ????????????????????????????????????????????????????????????
	//dbClient.Clauses(clause.OnConflict{
	//	UpdateAll: true,
	//}).Create(&user)

	// ?????????????????????SQL??????
	//dbClient.Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "id"}},
	//	DoUpdates: clause.Assignments(map[string]interface{}{"email": gorm.Expr("(select email from user1_tab limit 1)")}),
	//}).Create(&user)
}
