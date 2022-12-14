package model

import (
	"database/sql"
	"gorm.io/plugin/soft_delete"
	"time"
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

type UserTab struct {
	ID    int64  `json:"id" gorm:"column:id"`
	Name  string `json:"name" gorm:"column:name"`
	Age   uint   `json:"age" gorm:"column:age"`
	Email string `json:"email" gorm:"column:email"`
	//IsDel uint8  `json:"is_del" gorm:"column:is_del"`
	IsDel soft_delete.DeletedAt `json:"is_del" gorm:"softDelete:flag"`

	CreateTime uint `json:"create_time" gorm:"column:create_time"`
	UpdateTime uint `json:"update_time" gorm:"column:update_time"`
	//DeleteAt     gorm.DeletedAt `json:"delete_at" gorm:"column:delete_at"`
	Birthday     *time.Time     `json:"birthday" gorm:"column:birthday"`
	MemberNumber sql.NullString `json:"member_number" gorm:"column:member_number"`
}

func (u *UserTab) TableName() string {
	return "user_tab"
}

//func (u *UserTab) BeforeCreate(tx *gorm.DB) (err error) {
//	md5 := crypto.MD5.New()
//	md5.Write([]byte(u.Email))
//	u.Email = hex.EncodeToString(md5.Sum(nil))
//	u.IsDel = 0
//	return
//}

//func (u *UserTab) BeforeUpdate(tx *gorm.DB) (err error) {
//	// ??????????????????????????????????????????
//	if tx.Statement.Changed("Email") {
//		return errors.New("email not allowed to change")
//	}
//
//	// ????????????????????????????????????????????????
//	if tx.Statement.Changed() {
//		tx.Statement.SetColumn("UpdateTime", time.Now().Unix())
//	}
//
//	return nil
//}

//func (u *UserTab) BeforeDelete(tx *gorm.DB) (err error) {
//	if u.Email == "lzc" {
//		return errors.New("user not allowed to delete")
//	}
//	return
//}
