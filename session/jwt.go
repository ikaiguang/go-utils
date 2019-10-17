package gosession

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

// JwtType token type
// var customType JwtTokenType = "custom_type"
type JwtTokenType string

// token type
const (
	JwtTypeBackground JwtTokenType = "background" // background
	JwtTypeMember     JwtTokenType = "member"     // user
	JwtTypeVip        JwtTokenType = "vip"        // vip
)

// String string
func (j JwtTokenType) String() string {
	return string(j)
}

// GenerateToken token
// @param signParam *JwtSignParam
func GenerateToken(signParam *JwtSignParam) (string, error) {
	// token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, signParam.Claims)
	// token string
	tokenString, err := jwtToken.SignedString(signParam.User.signKey())
	if err != nil {
		return "", errors.WithStack(err)
	}

	// save user
	if err := signParam.saveUser(); err != nil {
		return tokenString, errors.WithStack(err)
	}
	return tokenString, nil
}

// DecodeToken decode token
// @param password 可选参数，用于解密 token
// @param password 不传此参数，则使用 new(JwtClaims).decodeKey()
func DecodeToken(tokenString string, password ...string) (*JwtSignParam, error) {
	var signParam = new(JwtSignParam)

	// token slice
	tokenSlice := strings.Split(tokenString, ".")
	if len(tokenSlice) < 2 {
		return nil, errors.WithStack(ErrTokenIncorrect)
	}

	// token payload
	var tokenPayload = new(JwtClaims)
	tokenJSONBytes, err := jwt.DecodeSegment(tokenSlice[1])
	if err != nil {
		return nil, errors.WithStack(ErrTokenIncorrect)
	}
	if err := json.Unmarshal(tokenJSONBytes, tokenPayload); err != nil {
		return nil, errors.WithStack(ErrTokenIncorrect)
	}

	// decode key
	var decodeKeyFn func(*jwt.Token) (interface{}, error)
	if len(password) > 0 {
		decodeKeyFn = func(*jwt.Token) (interface{}, error) { return []byte(password[0]), nil }
	} else {
		decodeKeyFn, err = tokenPayload.decodeKey()
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	// token
	var decodePayload = new(JwtClaims)
	jwtToken, err := jwt.ParseWithClaims(tokenString, decodePayload, decodeKeyFn)
	if err != nil {
		return nil, errors.WithStack(ErrTokenInvalid)
	}
	signParam.Claims = decodePayload

	// invalid
	if !jwtToken.Valid {
		return nil, errors.WithStack(ErrTokenInvalid)
	}
	return signParam, nil
}

// JwtClaims jwt.StandardClaims
type JwtClaims struct {
	Uid   int64        `json:"u_id,omitempty"`   // user id
	Uuid  string       `json:"uuid,omitempty"`   // user uuid
	UType JwtTokenType `json:"u_type,omitempty"` // user type
	jwt.StandardClaims                           // jwt
}

// RemoveSession remove session
func (j *JwtClaims) RemoveSession() error {
	if err := storage.RemoveUser(j); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// storageKey mysql/redis/... key
func (j *JwtClaims) storageKey() string {
	return j.UType.String() + "_" + j.Uuid + "_" + strconv.FormatInt(j.Uid, 10)
}

// decodeKey decode key
func (j *JwtClaims) decodeKey() (func(*jwt.Token) (interface{}, error), error) {
	// user
	user, err := storage.GetUser(j)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return func(*jwt.Token) (interface{}, error) {
		return []byte(user.signKey()), nil
	}, nil
}

// expireDuration expire
func (j *JwtClaims) expireDuration() time.Duration {
	second := j.ExpiresAt - j.NotBefore
	if j.NotBefore <= 0 {
		second = j.ExpiresAt - time.Now().Unix()
	}
	if second <= 0 {
		return 0
	}
	return time.Duration(second) * time.Second
}

// JwtUser user payload
// edit password -> clear access token(mysql/redis/...)
type JwtUser struct {
	Uid          int64        `json:"id,omitempty"`             // user id
	Uuid         string       `json:"uuid,omitempty"`           // user uuid
	Nickname     string       `json:"nickname,omitempty"`       // nickname
	Gender       int32        `json:"gender,omitempty"`         // 0-unknown,1-man,2-woman
	Avatar       string       `json:"avatar,omitempty"`         // avatar
	InviteCode   string       `json:"invite_code,omitempty"`    // invite code
	IsVip        bool         `json:"is_vip,omitempty"`         // vip
	UType        JwtTokenType `json:"u_type,omitempty"`         // user type
	SignTokenKey string       `json:"sign_token_key,omitempty"` // token key
}

// signKey sign key
func (j *JwtUser) signKey() []byte {
	return []byte(j.SignTokenKey)
}

// JwtSignParam jwt sign param
type JwtSignParam struct {
	User   *JwtUser   // background/member/vip/...
	Claims *JwtClaims // claims
}

// saveKey save key
func (j *JwtSignParam) saveUser() error {
	if err := storage.SaveUser(j); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
