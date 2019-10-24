package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gzlj/client-go-demo/pkg/handler"
)

type APIServer struct {
	Engine *gin.Engine
}

func (s *APIServer) RegistryApi() {
	registryK8sHandler(s.Engine)
}

func registryK8sHandler(engine *gin.Engine) {
	engine.GET("/pod/count", handler.NamespacedPodCounts)
	engine.POST("/namespace", handler.CreateNamespace)
	engine.POST("/workload", handler.CreateWorkload)
}
