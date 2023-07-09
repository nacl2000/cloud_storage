package dal

import (
	"github.com/nacl2000/clould_storage/model"
)

func QueryPasswdByMd5(uname string) string {
	var result model.User
	res := GetModelDB(model.User{}).First("username = ?", uname).Scan(&result)
	if res.Error != nil {

	}
	if result.Username != uname {
		return ""
	}
	return result.Password
}

func CreateUser(userInfo model.User) {
	res := GetDBHandler().Create(userInfo)
	if res.Error != nil {
		// todo
	}
}
