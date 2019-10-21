package auth

import "golang.org/x/crypto/bcrypt"

// Encrypt encrypts the plain text with bcrypt.
//加密密码
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

//比较如果相同，则将加密的文本与纯文本进行比较。
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
