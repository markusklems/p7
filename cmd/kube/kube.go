package main

import (
	"flag"
	"fmt"
	"log"

	"k8s.io/client-go/1.5/kubernetes"
	"k8s.io/client-go/1.5/pkg/api"
	"k8s.io/client-go/1.5/rest"
	"k8s.io/client-go/1.5/tools/clientcmd"
)

var (
	kubeconfig = flag.String("kubeconfig", "/etc/kubernetes/admin.conf", "absolute path to the kubeconfig file")
)

type kubeContext struct {
	client *kubernetes.Clientset
}

type job struct {
	Name string
	Uid  string
}

func newKubeContext(inCluster bool) *kubeContext {
	var config *rest.Config
	var err error
	if inCluster {
		// creates the in-cluster config
		config, err = rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
	} else {
		// uses the current context in kubeconfig
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err.Error())
		}
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return &kubeContext{
		client: clientset,
	}
}

func (k kubeContext) listJobs() {
	jobs, err := k.client.Batch().Jobs("").List(api.ListOptions{})
	if err != nil {
		log.Printf("Couldn't find any jobs: %q", err)
	}
	for _, job := range jobs.Items {
		fmt.Printf("Jobame: %s\n", job.GetName())
		fmt.Printf("Job cluster name: %s\n", job.GetClusterName())
		fmt.Printf("JobUID: %s\n", job.GetUID())
	}

}
