package goredis

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

func NewClient() (*redis.Client, error) {
	rClient := redis.NewClient(InitConfig())

	// ping
	if err := rClient.Ping().Err(); err != nil {
		return nil, errors.WithStack(err)
	}
	return rClient, nil
}
