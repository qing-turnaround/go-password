package password

import (
	"crypto/md5"
	"encoding/hex"
)

// encryptPassword 将密码字符串加密
func EncryptPassword(password, salt string) string {
	h := md5.New()
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
// password：未进行处理的原密码，encryptPassword：进行过加密的密码，salt:加盐值
func verifyPassword(password, encryptPassword,salt string) bool {
	return EncryptPassword(password, salt) == encryptPassword
}
