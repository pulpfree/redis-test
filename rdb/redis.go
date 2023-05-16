package rdb

import (
	"context"
	"fmt"
	"os"

	"github.com/nitishm/go-rejson/v4"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

type RedisDb struct {
	client     *redis.Client
	connectUri string
	JH         *rejson.Handler
	stage      string
}

func New(stage string) (rdb *RedisDb, err error) {

	rdb = &RedisDb{
		stage: stage,
	}
	jsonh := rejson.NewReJSONHandler()

	rdb.setConnectUri()
	opt, err := redis.ParseURL(rdb.connectUri)
	if err != nil {
		log.Fatalf("failed redis connect %v", err)
		return nil, err
	}

	log.Info("Connecting to redis...")
	client := redis.NewClient(opt)
	jsonh.SetGoRedisClientWithContext(context.Background(), client)

	return &RedisDb{
		client: client,
		JH:     jsonh,
	}, nil
}

//
// ========================== Exported Methods ================================================= //
//

//
// ========================== private methods ================================================== //
//

func (r *RedisDb) setConnectUri() error {

	connectUri := os.Getenv("REDIS_ENDPOINT")
	if connectUri == "" {
		errStr := "REDIS_ENDPOINT not set"
		log.Fatal(errStr)
		return fmt.Errorf(errStr)
	}
	// currently we only have one stage, but switch setup to add additional
	switch r.stage {
	default:
		r.connectUri = connectUri
	}
	return nil
}
