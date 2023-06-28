package calculte

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Sum(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func GetSaltingPwd(passwd string) string {

	return ""
}
