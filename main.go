package main

import (
	"context"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"github.com/openshift/library-go/pkg/config/helpers"
	configv1 "github.com/openshift/api/config/v1"
)

func main() {
	config, err := helpers.GetKubeConfigOrInClusterConfig("", configv1.ClientConnectionOverrides{})
	if err != nil {
		panic(err.Error())
	}
	config.Timeout = 30 * time.Second
	config.QPS = 3
	config.Burst = 5

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	for {
		_, err := clientset.CoreV1().Secrets("test-01").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			fmt.Println(fmt.Sprintf("error while listing secrets, err = %v", err))
		}
	}
}
