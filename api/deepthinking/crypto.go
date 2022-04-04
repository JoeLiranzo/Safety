package deepthinking

import (
	"crypto/md5"
	"encoding/hex"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

func Encrypt(text *string) {
	md5 := md5.New()
	bi := big.NewInt(0)
	md5.Write([]byte(*text))
	hexstr := hex.EncodeToString(md5.Sum(nil))
	bi.SetString(hexstr, 16)
	*text = bi.String()
}

func Decrypt(text *string) {

}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
