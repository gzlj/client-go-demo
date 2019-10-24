package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gzlj/client-go-demo/pkg/client"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func CreateNamespace(ctx *gin.Context) {
	var (
		name      string
		err       error
		namespace corev1.Namespace
	)
	name = ctx.Query("name")
	if name == "" {
		ctx.JSON(400, "please input a name for namespace to be created.")
		return
	}

	namespace = corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Status: corev1.NamespaceStatus{
			Phase: corev1.NamespaceActive,
		},
	}
	_, err = client.G_ClientSet.CoreV1().Namespaces().Create(&namespace)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, gin.H{
		"namespace": name,
		"created":   true,
	})

}
