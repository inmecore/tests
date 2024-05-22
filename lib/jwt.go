package lib

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/sync/singleflight"
	"time"
)

const (
	Issuer      = "Tester"
	SigningKey  = "TesterSigningKey"
	BufferTime  = 24 * time.Hour
	ExpiredTime = 7 * 24 * time.Hour
)

type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	BufferTime int64
	BaseClaims
	jwt.RegisteredClaims
}

type BaseClaims struct {
	Uid int
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("can't handle this token")
)

func NewJWT() *JWT {
	return &JWT{SigningKey: []byte(SigningKey)}
}

func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	return CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(BufferTime / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"GVA"},                         // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),       // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ExpiredTime)), // 过期时间 7天  配置文件
			Issuer:    Issuer,                                          // 签名的发行者
		},
	}
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims CustomClaims) (string, error) {
	v, err, _ := (&singleflight.Group{}).Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (i interface{}, e error) {
			return j.SigningKey, nil
		},
	)
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token == nil {
		return nil, TokenInvalid
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, TokenInvalid
}
