package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/idzharbae/grpc-stream-interceptor/interceptor/stream"
	"github.com/idzharbae/grpc-stream-interceptor/interceptor/unary"
	"github.com/idzharbae/grpc-stream-interceptor/stream/proto"
	"github.com/idzharbae/grpc-stream-interceptor/stream/server/datastructure"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/rand"
	"net"
	"sort"
)

type Server struct{}

func (s *Server) GetRandomNumber(ctx context.Context, req *proto.NumberRange) (*proto.RandomNumber, error) {
	num := rand.Int31n(req.GetMax()-req.GetMin()) + req.GetMin()
	return &proto.RandomNumber{Num: num}, nil
}

func (s *Server) SortInteger(stream proto.StreamService_SortIntegerServer) error {
	arr := make([]int, 0)
	for {
		num, err := stream.Recv()
		if err == io.EOF {
			sort.Ints(arr)
			resp := make([]int32, len(arr))
			for i := range arr {
				resp[i] = int32(arr[i])
			}
			return stream.SendAndClose(&proto.SortIntegerResp{
				SortedNumbers: resp,
			})
		}
		arr = append(arr, int(num.GetNumber()))
	}
}
func (s *Server) SortIntegerInteractive(stream proto.StreamService_SortIntegerInteractiveServer) error {
	var bst *datastructure.BST
	for {
		num, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		bst = datastructure.Insert(bst, num.GetNumber())
		if err := stream.Send(&proto.SortIntegerResp{SortedNumbers: datastructure.ToArray(bst)}); err != nil {
			return err
		}
	}
}

func (s *Server) PrintFibonacci(req *proto.FibonacciReq, stream proto.StreamService_PrintFibonacciServer) error {
	offset := req.GetOffset()
	limit := req.GetLimit()

	if offset < 0 || limit < 0 {
		return errors.New("offset and limit should not be negative")
	}

	prev := int32(0)
	current := int32(1)
	printed := int32(0)

	for idx := int32(1); printed < limit; idx++ {

		if idx > offset {
			if err := stream.Send(&proto.FibonacciResp{Num: current}); err != nil {
				return err
			}
			printed++
		}
		prev, current = current, current+prev
	}

	return nil
}

func main() {
	s := getServer()
	runServer(s, ":50005")
}

func runServer(s *grpc.Server, port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("listening to port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getServer() *grpc.Server {
	var grpcOptions []grpc.ServerOption
	grpcOptions = append(grpcOptions,
		grpc.UnaryInterceptor(unary.ServerLogger()),
		grpc.StreamInterceptor(stream.ServerLogger()),
	)

	s := grpc.NewServer(grpcOptions...)
	proto.RegisterStreamServiceServer(s, &Server{})
	return s
}
