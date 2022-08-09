package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func main() {
	fun1()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func createclient() *redis.Client {
	// 连接redis
	// 设置一些连接参数
	// 自带连接池
	opt := &redis.Options{}
	opt.Addr = "localhost:6379"
	opt.Password = "123456"
	opt.PoolSize = 10
	client := redis.NewClient(opt)
	_, err := client.Ping(context.TODO()).Result()
	must(err)
	return client
}

// 基础使用
func fun1() {
	ctx := context.TODO()
	var err error
	var val string
	var key string

	client := createclient()

	// 删除所有key
	keys, err := client.Keys(ctx, "test*").Result()
	must(err)
	for _, v := range keys {
		_, err = client.Del(ctx, v).Result()
		must(err)
	}

	// 字符串操作
	// 设置 tanght="abc" 并读取
	key = "mystring"
	err = client.Del(ctx, key).Err()
	must(err)
	err = client.Set(ctx, key, "abc", 0).Err()
	must(err)
	val, err = client.Get(ctx, key).Result()
	must(err)
	fmt.Println(val)

	// 设置 tanght=10086 并读取
	// redis 没有 int 类型, 用string来表示int
	// 然后 Get 到的结果需要自己转换成 int
	err = client.Del(ctx, key).Err()
	must(err)
	err = client.Set(ctx, key, 10086, 0).Err()
	must(err)
	val, err = client.Get(ctx, key).Result()
	must(err)
	ival, err := strconv.Atoi(val)
	must(err)
	fmt.Println(ival)

	// Hash操作
	key = "myhash"
	err = client.Del(ctx, key).Err()
	must(err)
	client.HSet(ctx, key, "name", "zhangsan", "age", 18)
	client.HSet(ctx, key, "school", "tju")
	hashdata, err := client.HGetAll(ctx, key).Result() // GetAll返回的是map[string]string
	must(err)
	fmt.Println(hashdata)

	err = client.Del(ctx, key).Err()
	must(err)
	client.HSet(ctx, key, map[string]interface{}{"name": "lisi", "age": 20, "school": "aaa"})
	hashdata, err = client.HGetAll(ctx, key).Result()
	must(err)
	fmt.Println(hashdata)

	// 列表操作
	// [head, 1, 1, 1, 1, 1, tail]
	// 左边是头, 右边是尾
	key = "mylist"
	err = client.Del(ctx, key).Err()
	must(err)
	client.LPush(ctx, key, "a") // [a]
	client.LPush(ctx, key, "b") // [b, a]
	client.LPush(ctx, key, "c") // [c, b, a]
	client.LPush(ctx, key, "d") // [d, c, b, a]
	client.LPush(ctx, key, "e") // [e, d, c, b, a]
	listdata, err := client.LRange(ctx, key, 0, 3).Result()
	must(err)
	fmt.Println(listdata) // [e, d, c, b]

	err = client.Del(ctx, key).Err()
	must(err)
	client.RPush(ctx, key, "a") // [a]
	client.RPush(ctx, key, "b") // [a, b]
	client.RPush(ctx, key, "c") // [a, b, c]
	client.RPush(ctx, key, "d") // [a, b, c, d]
	client.RPush(ctx, key, "e") // [a, b, c, d, e]
	listdata, err = client.LRange(ctx, key, 0, 3).Result()
	must(err)
	fmt.Println(listdata) // [a, b, c, d]

	val, err = client.LIndex(ctx, key, 0).Result()
	must(err)
	fmt.Println(val)
	val, err = client.LIndex(ctx, key, 1).Result()
	must(err)
	fmt.Println(val)
	val, err = client.LIndex(ctx, key, 2).Result()
	must(err)
	fmt.Println(val)

	val, err = client.LIndex(ctx, key, -1).Result()
	must(err)
	fmt.Println(val)
	val, err = client.LIndex(ctx, key, -2).Result()
	must(err)
	fmt.Println(val)
	val, err = client.LIndex(ctx, key, -3).Result()
	must(err)
	fmt.Println(val)
}

func luascript() {
	ctx := context.TODO()
	client := createclient()
	keys := []string{"k1", "k2"}
	args := []string{"a1", "a2"}
	luas := `return {KEYS[1],KEYS[2],ARGV[1],ARGV[2]}`
	hash := sha1.New()
	hash.Sum()
	hex.EncodeToString(sha1.Sum([]byte(luas)))
	val, err := client.Eval(ctx, luas, keys, args).Result() // 执行Lua脚本
	must(err)
	fmt.Println(val)
	client.ScriptLoad()   // 将Lua脚本缓存到redis服务器
	client.EvalSha()      // 通过脚本的sha来执行redis服务器中的Lua脚本
	client.ScriptKill()   // 杀死当前正在执行的Lua脚本
	client.ScriptFlush()  // 删除redis服务器缓存的所有Lua脚本
	client.ScriptExists() // 判断redis服务器是否存在此Lua脚本

	script := redis.NewScript("")
	// 将脚本缓存到redis服务器
	script.Load()
	// 执行redis服务器中缓存的脚本
	// Load的时候已经将脚本缓存到redis服务器了
	// EvalSha只需将脚本的sha和参数发送给redis就行了
	script.EvalSha()

	// 将整个脚本发送发给redis服务器并执行
	script.Eval()
	script.Hash()
}
