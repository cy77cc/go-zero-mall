package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go func() {
		<-time.After(3 * time.Second)
		fmt.Println("xxxxx")
	}()

	select {
	case <-ctx.Done():
		fmt.Println("超时了，不玩了")
	}

	append()

}
