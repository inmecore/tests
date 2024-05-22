package lib

import "time"

const (
	DefaultPasswordVal  string = "123456" // 默认密码
	DefaultPasswordHash string = "FF1757E51EC5E7F5D719B4623A4946BB"
	KeyEncode           string = "CC9Z28P04DPN455A5AD9CA59"
	KeyXClaims          string = "x-claims"
	KeyXToken           string = "x-token"
	KeyXUId             string = "x-uid"
)

var (
	Now = time.Date(2024, 1, 2, 3, 0, 0, 0, time.UTC)
)
