package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gzlj/client-go-demo/pkg/api"
)

func main() {

	server := &api.APIServer{
		Engine: gin.Default(),
	}
	server.RegistryApi()
	server.Engine.Run(":18080")

}
