package main

import (
	"log"
	"os"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		log.Fatal("no resource is mentioned \nExample command : ls pods")

	}

	if argsWithoutProg[0] != "pods" {
		log.Fatal("unknown resource kind")
	}

	var cs *kubernetes.Clientset
	cs.CoreV1().Pods("kube-system").List(meta_v1.ListOptions{})
}
