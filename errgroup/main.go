package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

// 用 errgroup 中的 Go 方法 替代 go 关键字 来执行协程
// errgroup 比 sync.WaitGroup 方便, 不需要 手动 Add 和 Done, 并且Wait可以获得error
// 只要某一个协程出错, 则 errgroup 的 ctx 自动结束
// 协程全部结束之后, Wait 才返回

func main() {
	fmt.Println(uuid.ClockSequence())
	fmt.Println(uuid.NewUUID())
	return
	eg, ctx := errgroup.WithContext(context.Background())
	for i := 0; i < 100; i++ {
		eg.Go(func() error {
			if rand.Intn(10) == 0 {
				fmt.Println("aaa")
				return fmt.Errorf("123123123")
			}
			time.Sleep(5 * time.Second)
			fmt.Println("bbb")
			return nil
		})
	}
	<-ctx.Done()
	fmt.Println("errgroup结束了")
	if err := eg.Wait(); err != nil {
		fmt.Printf("ERROR %v\n", err)
	}

	time.Sleep(100 * time.Second)
}
