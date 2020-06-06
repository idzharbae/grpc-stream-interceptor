package main

import (
	"context"
	"fmt"
	"github.com/idzharbae/grpc-stream-interceptor/stream/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/rand"
	"time"
)

func printFibonacci(client proto.StreamServiceClient) {
	req := &proto.FibonacciReq{Limit: 10, Offset: 10}
	stream, err := client.PrintFibonacci(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("fibonacci (%d, %d): ", 10, 10)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d ", resp.GetNum())
	}
	fmt.Println()
}

func sortInteger(client proto.StreamServiceClient) {
	stream, err := client.SortInteger(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		if err := stream.Send(&proto.SortIntegerReq{Number: rand.Int31() % 100}); err != nil {
			log.Fatal(err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetSortedNumbers())
}

func sortIntegerInteractive(client proto.StreamServiceClient) {
	stream, err := client.SortIntegerInteractive(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan bool)

	// write stream
	go func() {
		for i := 0; i < 10; i++ {
			num := rand.Int31() % 100
			fmt.Printf("send input: %d\n", num)
			if err := stream.Send(&proto.SortIntegerReq{Number: num}); err != nil {
				log.Fatal(err)
			}
		}
		if err := stream.CloseSend(); err != nil {
			log.Fatal(err)
		}
	}()

	// read stream
	go func() {
		for {
			reply, err := stream.Recv()
			if err == io.EOF {
				done <- true
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("sorted array: %v\n", reply.GetSortedNumbers())
		}
	}()

	<-done
}

func main() {
	rand.Seed(time.Now().Unix())

	conn, err := grpc.Dial(":50005", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := proto.NewStreamServiceClient(conn)

	// print fibonacci
	printFibonacci(client)

	// sort integer
	sortInteger(client)

	// sort integer interactive
	sortIntegerInteractive(client)
}
