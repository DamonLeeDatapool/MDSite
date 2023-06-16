package funcs

import (

	//"os"
	//"encoding/binary"
	//"strconv"
	//"log"
	//"Errors"
	//"MServer/globalVar"
	"golang.org/x/crypto/bcrypt"
)

// 三元運算子
func AorB(condition bool, a, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

// 密碼加密
func PasswordHash(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

// 密碼比對
func PasswordVerify(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
