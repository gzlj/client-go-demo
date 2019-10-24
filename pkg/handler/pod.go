package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gzlj/client-go-demo/pkg/client"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func NamespacedPodCounts(ctx *gin.Context) {
	var (
		namespace string
		podList   *corev1.PodList
		err       error
		count     int
	)
	namespace = ctx.DefaultQuery("namespace", "default")

	podList, err = client.G_ClientSet.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(500, nil)
		return
	}
	count = len(podList.Items)
	ctx.JSON(200, gin.H{
		"namespace": namespace,
		"count":     count,
	})

}
