package services

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

type Redis struct {
	user string
	pass string
	host string
	port string
	uri  string
}

func InitRedis(user string, pass string, host string, port string) Redis {
	uri := fmt.Sprintf("rediss://%s:%s@%s:%s", user, pass, host, port)
	return Redis{user: user, pass: pass, host: host, port: port, uri: uri}
}

func (r *Redis) Connect() (*redis.Client, error) {
	opt, err := redis.ParseURL(r.uri)
	if err != nil {
		log.Fatalln(err)
	}

	// Connect to Redis
	client := redis.NewClient(opt)
	fmt.Println("Connected to Redis!")

	return client, nil
}
