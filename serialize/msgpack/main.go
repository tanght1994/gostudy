package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/msgpack/v5"
)

func main() {
	type Item struct {
		Foo string
	}

	b, err := msgpack.Marshal(&Item{Foo: "bar"})
	if err != nil {
		panic(err)
	}

	var item Item
	err = msgpack.Unmarshal(b, &item)
	if err != nil {
		panic(err)
	}
	fmt.Println(item.Foo)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func createRedisClient() *redis.Client {
	opt := &redis.Options{}
	opt.Addr = "localhost:6379"
	// opt.Password = "123456"
	opt.PoolSize = 10
	client := redis.NewClient(opt)
	_, err := client.Ping(context.TODO()).Result()
	must(err)
	return client
}

type DelayQueue struct {
	topic string
}

func (this *DelayQueue) Add(exectime time.Time) error {
	time.Now()
	return nil
}

func (this *DelayQueue) run(ctx context.Context, rc *redis.Client) {
	lua := `
local data = redis.call("ZRANGEBYSCORE", KEYS[1], 0, redis.call("TIME")[1], "LIMIT", 0, 1000)
if(#data ~= 0)
then
	redis.call("RPUSH", KEYS[2], unpack(data))
end
redis.call("ZREMRANGEBYRANK", KEYS[1], 0, #data-1)
return #data
	`
	for {
		rc.Eval(ctx, lua, []string{"delay_queue", "ready_queue"}).Result()
	}
}
