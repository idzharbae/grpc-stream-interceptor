package unary

import (
	"context"
	"google.golang.org/grpc"
	"os"
	"sync"
	"time"
)

var file *os.File
var once sync.Once

func ServerLogger() grpc.UnaryServerInterceptor {
	once.Do(func() {
		logFile, err := os.Create("./unarylog.txt")
		if err != nil {
			panic(err)
		}
		file = logFile
	})
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {

		err := writeLog(info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func writeLog(methodString string) error {
	data := time.Now().Format("[2006-01-02 15:04:05]") + " - " + methodString + "\n"
	_, err := file.Write([]byte(data))
	if err != nil {
		return err
	}
	return nil
}
