// internal/infrastructure/http/http_server.go
package http

import (
    "github.com/gin-gonic/gin"
    "github.com/nuhmanudheent/go-microservices/internal/interfaces/http"
    "github.com/nuhmanudheent/go-microservices/internal/usecase"
)

func StartHTTPServer() {
    r := gin.Default()

    userUseCase := usecase.NewUserUseCase()
    userHandler := http.NewUserHandler(userUseCase)

    r.GET("/user/:id", userHandler.GetUser)

    r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
