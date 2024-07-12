// cmd/service/main.go
package main

import (
	"sync"

	grpcServer "github.com/nuhmanudheent/go-microservices/internal/infrastructure/grpc"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		grpcServer.StartGRPCServer()
	}()

	wg.Wait()
}
