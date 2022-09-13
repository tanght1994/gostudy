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
	binary_val()
	return
	luascript()
	fun1()
}

func aaaaaa() {
	ctx := context.Background()
	client := createclient()
	must(client.Set(ctx, "tanght", []byte{200, 201, 202}, 0).Err())
	res := client.Get(ctx, "tanght")
	must(res.Err())
	val, err := res.Result()
	must(err)
	fmt.Println([]byte(val))
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
	// opt.Password = "123456"
	opt.PoolSize = 10
	client := redis.NewClient(opt)
	_, err := client.Ping(context.TODO()).Result()
	must(err)
	return client
}

func clusterclient() *redis.ClusterClient {
	opt := &redis.ClusterOptions{}
	opt.Addrs = []string{}
	client := redis.NewClusterClient(opt)
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

// lua脚本
func luascript() {
	ctx := context.TODO()
	client := createclient()
	keys := []string{"k1", "k2"}
	args := []string{"a1", "a2"}
	luas := `return {KEYS[1],KEYS[2],ARGV[1],ARGV[2]}`
	hash := sha1.Sum([]byte(luas))
	luashash := hex.EncodeToString(hash[:])

	// 删除redis服务器缓存的所有Lua脚本
	tmp1, err := client.ScriptFlush(ctx).Result()
	must(err)
	fmt.Println(tmp1)

	// 判断redis服务器是否存在此Lua脚本
	tmp5, err := client.ScriptExists(ctx, luashash).Result()
	must(err)
	fmt.Println(tmp5)

	// 执行Lua脚本
	// 将整个lua脚本发送给redis服务器并执行
	tmp2, err := client.Eval(ctx, luas, keys, args).Result()
	must(err)
	fmt.Println(tmp2)

	// 判断redis服务器是否存在此Lua脚本
	tmp6, err := client.ScriptExists(ctx, luashash).Result()
	must(err)
	fmt.Println(tmp6)

	// 将Lua脚本缓存到redis服务器
	// 其实上一步Eval的时候, 脚本已经被redis服务器缓存了
	// tmp3是此脚本的hash
	tmp3, err := client.ScriptLoad(ctx, luas).Result()
	must(err)
	fmt.Println(tmp3)

	// 判断redis服务器是否存在此Lua脚本
	tmp7, err := client.ScriptExists(ctx, luashash).Result()
	must(err)
	fmt.Println(tmp7)

	// 通过脚本的sha来执行redis服务器中的Lua脚本
	tmp4, err := client.EvalSha(ctx, luashash, keys, args).Result()
	must(err)
	fmt.Println(tmp4)
}

// 使用集群
func cluster() {
	client := clusterclient()
	// ClusterClient会自动根据key来选择合适的redis节点
	// 省去了redis服务帮我们重定向的操作, 提高了操作效率
	client.Get(context.TODO(), "tanght").Result()
}

// redis存取二进制数据
func binary_val() {
	client := createclient()
	ctx := context.Background()
	key := "tanght"
	defer client.Del(ctx, key)

	fmt.Println("Set Get 二进制数据测试-------------")

	// Set 可以存入二进制的val，比如[]byte
	// Get 可以获取二进制的val
	must(client.Del(ctx, key).Err())
	must(client.Set(ctx, key, []byte{200, 201}, 0).Err())
	val, _ := client.Get(ctx, key).Result()
	fmt.Println(val) //乱码 因为[200, 201]这个字节流不是ascii字符码
	fmt.Println([]byte(val))

	fmt.Println("Hset Hget 二进制数据测试-------------")

	// hset 设置 二进制数据
	// hget 获取 二进制数据
	must(client.Del(ctx, key).Err())
	must(client.HSet(ctx, key, "k1", "v1", "k2", []byte{200, 201}).Err())

	val, _ = client.HGet(ctx, key, "k1").Result()
	fmt.Println(val)

	val, _ = client.HGet(ctx, key, "k2").Result()
	fmt.Println(val) //乱码
	fmt.Println([]byte(val))

	fmt.Println("Hset Hget 二进制数据测试-------------")
	must(client.Del(ctx, key).Err())
	must(client.SAdd(ctx, key, "abc", "nihao", []byte{200, 201}, []byte{202, 203}).Err())

	// 插入 abc 因为已经存在了所以插入失败
	cnt, _ := client.SAdd(ctx, key, "abc").Result()
	fmt.Println(cnt) // 0

	// 插入 []byte{200, 201} 因为已经存在了所以插入失败
	cnt, _ = client.SAdd(ctx, key, []byte{200, 201}).Result()
	fmt.Println(cnt) // 0

	// 插入 []byte{201, 202} 因为集合中不存在所以插入成功
	cnt, _ = client.SAdd(ctx, key, []byte{201, 202}).Result()
	fmt.Println(cnt) // 1

	// 获取集合中所有元素
	ss, _ := client.SMembers(ctx, key).Result()
	for _, v := range ss {
		fmt.Println([]byte(v))
	}

	fmt.Println("LPush 二进制数据测试-------------")
	must(client.Del(ctx, key).Err())
	must(client.LPush(ctx, key, []byte{'a', 'b', 'c'}).Err())
	must(client.LPush(ctx, key, []byte{200, 201, 202}).Err())
	val1, _ := client.LPop(ctx, key).Result()
	val2, _ := client.LPop(ctx, key).Result()
	fmt.Println([]byte(val1))
	fmt.Println([]byte(val2))
}
