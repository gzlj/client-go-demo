package client

import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
)

var (
	G_ClientSet *kubernetes.Clientset
)

func init() {
	var (
		kubeConfigFile *string
		err            error
		config         *rest.Config
	)
	kubeConfigFile = flag.String("kubeConfigFile", "/root/.kube/config", "kubernetes config file path")
	flag.Parse()
	config, err = clientcmd.BuildConfigFromFlags("", *kubeConfigFile)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	G_ClientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("connect k8s success")
	}
}
