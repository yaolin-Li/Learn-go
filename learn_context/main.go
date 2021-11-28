package main

import (
	"context"
	"fmt"
	"time"
)

func doSometing(ctx context.Context) {
	select {
	case <- time.After(5 * time.Second): // 5秒过去了
		fmt.Println("finish do something")
	case <- ctx.Done():
		err := ctx.Err()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func main() {
	// 创建空context
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	//ctx, cancel := context.WithTimeout(ctx, 2 * time.Second) //两秒之后自动cancel, 一段时间
	//ctx, cancel := context.WithDeadline(ctx,具体时间 )

	go func() {
		time.Sleep(6 * time.Second)
		cancel()
	}()

	doSometing(ctx)


}
