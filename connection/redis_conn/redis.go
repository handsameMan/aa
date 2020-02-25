package redis_conn

import (
	"archie/utils"
	"archie/utils/configer"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func GetRedis() (redis.Conn, error) {
	return redis.Dial("tcp", getRedisURL())
}

func GetRedisConnMust(cb func(conn redis.Conn)) {
	conn, err := GetRedis()
	utils.Check(err)

	defer conn.Close()
	cb(conn)
}

func getRedisURL() string {
	redisConfig := configer.LoadRedisConfig()

	return fmt.Sprintf("%s:%s", redisConfig.Bind, redisConfig.Port)
}
