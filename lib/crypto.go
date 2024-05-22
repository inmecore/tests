package lib

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"github.com/forgoer/openssl"
	"strings"
)

func MD5(text string) string {
	m := md5.New()
	m.Write([]byte(text))
	return hex.EncodeToString(m.Sum(nil))
}

func Encode(src, key string) (string, error) {
	text, err := openssl.Des3ECBEncrypt([]byte(src), []byte(key), openssl.PKCS7_PADDING)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(text), nil
}

func Decode(src, key string) (string, error) {
	dst, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}
	text, err := openssl.Des3ECBDecrypt(dst, []byte(key), openssl.PKCS7_PADDING)
	if err != nil {
		return "", err
	}
	return string(text), nil
}

func Password(source string) string {
	return strings.ToUpper(MD5("ZF9756C254TSVICRYTQMNYU5" + source))
}
