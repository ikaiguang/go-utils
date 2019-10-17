package gosession

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	SetStorage(NewMemoryStorage())
	//SetStorage(NewFileStorage(filepath.Join(pwdPath, "_output/session")))
	//SetStorage(NewRedisStorage(nil, ""))
	var nowTime = time.Now()
	var user = &JwtUser{
		Uid:          123456789,
		Uuid:         "abc",
		Nickname:     "Nickname",
		Gender:       0,
		Avatar:       "https://uufff.com/abc.jpg",
		InviteCode:   "1234",
		IsVip:        true,
		UType:        JwtTypeMember,
		SignTokenKey: "key",
	}
	var claims = &JwtClaims{
		Uid:   user.Uid,
		Uuid:  user.Uuid,
		UType: user.UType,
		StandardClaims: jwt.StandardClaims{
			Audience:  user.Uuid,                      // aud 目标收件人(签发给谁)
			ExpiresAt: nowTime.Add(time.Hour).Unix(),  // exp 过期时间(有效期时间 exp)
			Id:        nowTime.Format(time.StampNano), // jti 编号
			IssuedAt:  nowTime.Unix(),                 // iat 签发时间
			Issuer:    "issuer",                       // iss 签发者
			NotBefore: nowTime.Unix(),                 // nbf 生效时间(nbf 时间后生效)
			Subject:   "test",                         // sub 主题
		},
	}
	var signParam = &JwtSignParam{
		User: user, Claims: claims,
	}
	tokenString, err := GenerateToken(signParam)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("token string : %s \n", tokenString)

	var decodeParam *JwtSignParam
	decodeParam, err = DecodeToken(tokenString, user.SignTokenKey)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("token string : %#v \n", decodeParam)
}
