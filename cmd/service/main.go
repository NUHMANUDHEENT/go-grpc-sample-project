// cmd/service/main.go
package main

import (
    "sync"
    grpcServer "github.com/nuhmanudheent/go-microservices/internal/infrastructure/grpc"
    httpServer "github.com/nuhmanudheent/go-microservices/internal/infrastructure/http"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        defer wg.Done()
        grpcServer.StartGRPCServer()
    }()

    go func() {
        defer wg.Done()
        httpServer.StartHTTPServer()
    }()

    wg.Wait()
}
