package gosession

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// SessionInterface interface
type SessionInterface interface {
	SaveUser(*JwtSignParam) error
	GetUser(*JwtClaims) (*JwtUser, error)
	RemoveUser(*JwtClaims) error
}

// storage storage
var storage SessionInterface = NewFileStorage(filepath.Join(pwdPath, "_output/session"))

// SetStorage storage
func SetStorage(s SessionInterface) {
	storage = s
}

// NewMemoryStorage memory storage
func NewMemoryStorage() *memoryStorage {
	return &memoryStorage{users: make(map[string]*JwtUser)}
}

// memoryStorage memory
type memoryStorage struct {
	mu    sync.Mutex
	users map[string]*JwtUser
}

// SaveUser save user
func (m *memoryStorage) SaveUser(signParam *JwtSignParam) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.users[signParam.Claims.storageKey()] = signParam.User

	return nil
}

// GetUser get user
func (m *memoryStorage) GetUser(claims *JwtClaims) (*JwtUser, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	user, ok := m.users[claims.storageKey()]
	if !ok {
		return nil, ErrTokenInvalid
	}
	return user, nil
}

// RemoveUser remove user
func (m *memoryStorage) RemoveUser(claims *JwtClaims) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := claims.storageKey()
	if _, ok := m.users[key]; ok {
		delete(m.users, key)
	}
	return nil
}

// NewRedisStorage redis
func NewRedisStorage(cc *redis.Client, keyPrefix string) *redisStorage {
	return &redisStorage{cc: cc, prefix: keyPrefix}
}

// redisStorage redis
type redisStorage struct {
	cc     *redis.Client // redis client
	prefix string        // key prefix
}

// key redis key
func (r *redisStorage) key(key string) string {
	return r.prefix + key
}

// SaveUser save user
func (r *redisStorage) SaveUser(signParam *JwtSignParam) error {
	// json
	userBytes, err := json.Marshal(signParam.User)
	if err != nil {
		return errors.WithStack(err)
	}

	// save
	cmd := r.cc.Set(r.key(signParam.Claims.storageKey()), userBytes, signParam.Claims.expireDuration())
	if err := cmd.Err(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// GetUser get user
func (r *redisStorage) GetUser(claims *JwtClaims) (*JwtUser, error) {
	// get user
	cmd := r.cc.Get(r.key(claims.storageKey()))
	if err := cmd.Err(); err != nil {
		if err == redis.Nil {
			return nil, ErrTokenInvalid
		} else {
			return nil, errors.WithStack(err)
		}
	}

	// user
	var user = new(JwtUser)
	if err := json.Unmarshal([]byte(cmd.Val()), user); err != nil {
		return nil, errors.WithStack(err)
	}
	return user, nil
}

// RemoveUser remove user
func (r *redisStorage) RemoveUser(claims *JwtClaims) error {
	if err := r.cc.Del(r.key(claims.storageKey())).Err(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// NewFileStorage file
func NewFileStorage(path string) *fileStorage {
	return &fileStorage{path: path}
}

// fileStorage file
// @param path rel path => $pwd/path
// @param path abs path =>  path
type fileStorage struct {
	path string // path
}

// file
var (
	fileStorageMutex sync.Once
)

// absPath abs path
func (f *fileStorage) absPath() string {
	if filepath.IsAbs(f.path) {
		return f.path
	}
	return filepath.Join(pwdPath, f.path)
}

// SaveUser save user
func (f *fileStorage) SaveUser(signParam *JwtSignParam) error {
	// file
	absPath := f.absPath()

	// mkdir
	var err error
	fileStorageMutex.Do(func() {
		err = os.MkdirAll(absPath, 0744)
	})
	if err != nil {
		return errors.WithStack(err)
	}

	// open file
	filename := filepath.Join(absPath, signParam.Claims.storageKey())
	sessionFile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return errors.WithStack(err)
	}
	defer sessionFile.Close()

	// marshal json
	userBytes, err := json.Marshal(signParam.User)
	if err != nil {
		return errors.WithStack(err)
	}

	// write file
	if _, err := sessionFile.Write(userBytes); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// GetUser get user
func (f *fileStorage) GetUser(claims *JwtClaims) (*JwtUser, error) {
	// open file
	filename := filepath.Join(f.absPath(), claims.storageKey())
	userBytes, err := ioutil.ReadFile(filename)
	if os.IsNotExist(err) {
		return nil, ErrTokenInvalid
	}

	// user
	var user = new(JwtUser)
	if err := json.Unmarshal(userBytes, user); err != nil {
		return nil, errors.WithStack(err)
	}
	return user, nil
}

// RemoveUser remove user
func (f *fileStorage) RemoveUser(claims *JwtClaims) error {
	filename := filepath.Join(f.absPath(), claims.storageKey())
	if err := os.Remove(filename); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return errors.WithStack(err)
	}
	return nil
}
