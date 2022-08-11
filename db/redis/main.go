package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	_ "github.com/go-redis/redis/v8"
	"github.com/gomodule/redigo/redis"
)

func main() {
	f3()
}

func f2() error {
	keys := []string{}
	for i := 0; i < 10000; i++ {
		keys = append(keys, fmt.Sprintf("key%d", i))
	}

	cn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Printf("redis.Dial error %v", err)
		return err
	}

	key := ""
	t0 := time.Now()
	for i := 0; i < 10000*10; i++ {
		key = keys[rand.Int()%len(keys)]
		cn.Do("SET", key, key)
	}
	t1 := time.Now()
	fmt.Println(t1.Sub(t0))

	return nil
}

func f3() error {
	keys := []string{}
	for i := 0; i < 10000; i++ {
		keys = append(keys, fmt.Sprintf("key%d", i))
	}
	pool := redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
		MaxActive: 100,
		Wait:      true,
	}
	start := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(10000 * 10)
	for i := 0; i < 10000*10; i++ {
		go func() {
			defer wg.Done()
			<-start
			cn := pool.Get()
			key := keys[rand.Int()%len(keys)]
			_, err := cn.Do("SET", key, key)
			if err != nil {
				fmt.Println("cn.Do error", err)
			}
			// fmt.Println(res)
			if cn.Close() != nil {
				fmt.Println("err")
			}
		}()
	}
	fmt.Println("ready")
	t0 := time.Now()
	close(start)
	fmt.Println("go")
	wg.Wait()
	t1 := time.Now()
	fmt.Println(t1.Sub(t0))
	return nil
}

// func f1() {
// 	// 先创建几个DialOption玩儿玩儿
// 	// DialOption用于配置redis客户端
// 	// 比如设置客户端名字(命令为CLIENT SETNAME haha)
// 	// 使用密码登陆redis(命令为AUTH your_password)
// 	// 使用用户+密码登陆redis(命令为AUTH your_name your_password)
// 	// 等等...
// 	ocn := redis.DialClientName("tanghttest")
// 	opw := redis.DialPassword("123456")
// 	// Dial接收N个DialOption选项,上面创建的，这里给它传递进去就行了
// 	c, e := redis.Dial("tcp", "localhost:6379", ocn, opw)
// 	if e != nil {
// 		fmt.Println("连接redis失败", e.Error())
// 	}
// 	// 然后就可以使用c了，c是redis的连接，通过c给redis发命令就行了
// 	// redis返回的是interface类型，可以使用redis包提供的类型转换函数进行转换
// 	// 比如redis.String()将redis的返回值转换为string类型
// 	fmt.Println(redis.String(c.Do("GET", "tanght")))
// 	fmt.Println(redis.String(c.Do("SET", "tanght", "100")))
// 	fmt.Println(redis.String(c.Do("GET", "tanght")))
// }

func f4() {
	// p := redis.NewPool(func() (redis.Conn, error) {return redis.Dial("tcp", "localhost:6379")}, 1000)
}
