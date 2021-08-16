package distributed_locker

import "math/rand"

var randChoice = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890~!@#$%^&*()_+{}<>?,./"

// RandStr 生成随机字符串
// length 生成字符串的长度
func RandStr(length int) string {
	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = randChoice[rand.Intn(len(randChoice))]
	}
	return string(buf)
}
