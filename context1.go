package main

import (
	"context"
	"fmt"
)

// func worker(ctx context.Context) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("worker is stopped:", ctx.Err())
// 			return
// 		default:
// 			fmt.Println("working....")
// 			time.Sleep(2 * time.Second)
// 		}
// 	}

// }

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	go worker(ctx)
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("stopping worker forcefully")
// 	cancel()
// 	time.Sleep(2 * time.Second)

// }

func oddoreven(ctx context.Context, number int) string {
	select {
	case <-ctx.Done():
		return fmt.Sprintf("process cancelled ", ctx.Err())
	default:
		if number%2 == 0 {
			return fmt.Sprintf("number %d is even", number)
		}
		return fmt.Sprintf("number %d is even", number)
	}
}

func main() {
	ctx := context.TODO()
	ctx, cancel := context.WithCancel(ctx)
	data := oddoreven(ctx, 10)
	fmt.Println(data)
	fmt.Println("after canceling")
	cancel()
	data = oddoreven(ctx, 12)
	fmt.Println(data)
}
