package xcontext

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

func Rpc(ctx context.Context, url string) error {
	result := make(chan int)
	err := make(chan error)

	go func() {
		// 进行RPC调用，并且返回是否成功，成功通过result传递成功信息，错误通过error传递错误信息
		isSuccess := true

		if xrandom.RandRange(0, 2) > 0 {
			isSuccess = false
		}

		if isSuccess {
			result <- 1
		} else {
			err <- errors.New("some error happen")
		}
	}()
	fmt.Println("阻塞之前！")
	select {
	case <-ctx.Done():
		// 其他RPC调用调用失败
		fmt.Println("其他rpc 调用失败", ctx.Err())
		return ctx.Err()
	case e := <-err:
		// 本RPC调用失败，返回错误信息
		fmt.Println("本rpc 调用失败", url)
		return e
	case <-result:
		// 本RPC调用成功，不返回错误信息
		fmt.Println("本rpc 调用成功", url)
		return nil

	}
}

func mainWork() {
	ctx, cancel := context.WithCancel(context.Background())

	// RPC1调用
	err := Rpc(ctx, "http://rpc_1_url")
	if err != nil {
		return
	}

	wg := sync.WaitGroup{}

	// RPC2调用
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := Rpc(ctx, "http://rpc_2_url")
		if err != nil {
			cancel()
		}
	}()

	// RPC3调用
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := Rpc(ctx, "http://rpc_3_url")
		if err != nil {
			cancel()
		}
	}()

	// RPC4调用
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := Rpc(ctx, "http://rpc_4_url")
		if err != nil {
			cancel()
		}
	}()

	wg.Wait()

}
