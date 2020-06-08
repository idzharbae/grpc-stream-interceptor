package stream

import (
	"google.golang.org/grpc"
	"os"
	"sync"
	"time"
)

var file *os.File
var once sync.Once

func ServerLogger() grpc.StreamServerInterceptor {
	once.Do(func() {
		logFile, err := os.Create("./streamlog.txt")
		if err != nil {
			panic(err)
		}
		file = logFile
	})
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

		err := writeLog(info.FullMethod)
		if err != nil {
			return err
		}

		return handler(srv, ss)
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
