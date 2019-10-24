package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gzlj/client-go-demo/pkg/client"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

type WorkloadConfig struct {
	Type      string            `json:"type"`
	Name      string            `json:"name"`
	NameSpace string            `json:"nameSpace"`
	Image     string            `json:"image"`
	Replicas  *int32            `json:"replicas"`
	Labels    map[string]string `json:"labels"`
	App       string            `json:"app"`
}

func CreateWorkload(ctx *gin.Context) {
	var (
		config WorkloadConfig
		err    error
	)

	if err = ctx.ShouldBindJSON(&config); err != nil {
		ctx.JSON(400, "requet body is not correct.")
		return
	}
	fmt.Println("config after ctx.ShouldBindJSON(&config): ", config)

	switch config.Type {
	case "Deployment":
		err = CreateDeployment(config)
	case "StatefulSet":
	case "DaemonSet":
	default:
		err = errors.New("please specify workload type.")
	}
	if err != nil {
		log.Println(err.Error())
		ctx.String(500, err.Error())
		return
	}
	ctx.JSON(200, gin.H{
		"type":      config.Type,
		"namespace": config.Name,
		"name":      config.Name,
		"created":   true,
	})

}

func CreateDeployment(config WorkloadConfig) (err error) {
	var (
		deployment *appsv1.Deployment
	)
	deployment = config.ToDeployment()
	_, err = client.G_ClientSet.AppsV1().Deployments(config.NameSpace).Create(deployment)
	return
}

func (config *WorkloadConfig) ToDeployment() (deployment *appsv1.Deployment) {
	var (
		d *appsv1.Deployment
	)
	selector := &metav1.LabelSelector{MatchLabels: config.Labels}
	d = &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      config.Name,
			Namespace: config.NameSpace,
			Labels:    config.Labels,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: selector,
			Replicas: config.Replicas,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: config.Labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  config.Name,
							Image: config.Image,
						},
					},
				},
			},
		},
	}
	deployment = d
	return
}
