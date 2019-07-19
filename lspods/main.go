package main

import (
	"fmt"
	"log"
	"os"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		log.Fatal("no resource is mentioned \nExample command : ls pods")
	}

	if argsWithoutProg[0] != "pods" {
		log.Fatal("unknown resource kind")
	}

	homeDir := os.Getenv("HOME")
	kubeconfigPath := homeDir + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	pods, err := clientset.CoreV1().Pods("kube-system").List(meta_v1.ListOptions{})
	if err != nil {
		log.Fatalf("failed to gte pods %v", err)
	}
	for _, p := range pods.Items {
		fmt.Println(p.Name)
	}
}
