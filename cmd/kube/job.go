package main

import (
	"fmt"
	"log"

	"k8s.io/client-go/1.5/kubernetes"
	api "k8s.io/client-go/1.5/pkg/api"
	errors "k8s.io/client-go/1.5/pkg/api/errors"
	"k8s.io/client-go/1.5/pkg/api/resource"
	"k8s.io/client-go/1.5/pkg/api/unversioned"
	api_v1 "k8s.io/client-go/1.5/pkg/api/v1"
	v1 "k8s.io/client-go/1.5/pkg/apis/batch/v1"
)

const namespace string = "default"

// operation represents a Kubernetes operation.
type operation interface {
	Do(c *kubernetes.Clientset)
}

type versionOperation struct{}

type deployOperation struct {
	image string
	name  string
	port  int
}

func (op *deployOperation) Do(c *kubernetes.Clientset) {
	err := op.createJob(c)
	// if job already exists and can't be created, delete it
	switch {
	case !errors.IsAlreadyExists(err):
		op.deleteJob(c)
	default:
		logger.Println("Job created")
	}
}

func (op *deployOperation) createJob(c *kubernetes.Clientset) error {
	appName := op.name

	// Define job spec.

	jobSpec := &v1.Job{
		TypeMeta: unversioned.TypeMeta{
			Kind:       "Job",
			APIVersion: "v1",
		},
		ObjectMeta: api_v1.ObjectMeta{
			Name: appName,
		},
		Spec: v1.JobSpec{
			Template: api_v1.PodTemplateSpec{
				ObjectMeta: api_v1.ObjectMeta{
					Name:   appName,
					Labels: map[string]string{"app": appName},
				},
				Spec: api_v1.PodSpec{
					Containers: []api_v1.Container{
						api_v1.Container{
							Name:  op.name,
							Image: op.image,
							Resources: api_v1.ResourceRequirements{
								Limits: api_v1.ResourceList{
									api_v1.ResourceCPU:    resource.MustParse("100m"),
									api_v1.ResourceMemory: resource.MustParse("256Mi"),
								},
							},
							ImagePullPolicy: api_v1.PullIfNotPresent,
						},
					},
					RestartPolicy: api_v1.RestartPolicyNever,
					DNSPolicy:     api_v1.DNSClusterFirst,
				},
			},
		},
	}

	jobs := c.Batch().Jobs(namespace)
	job, err := jobs.Create(jobSpec)
	if err != nil {
		return fmt.Errorf("Error during job creation: %s", err)
	}
	log.Printf("Job: %q created.", job)

	return nil
}

func (op *deployOperation) deleteJob(c *kubernetes.Clientset) error {
	deleteSpec := &api.DeleteOptions{}
	jobs := c.Batch().Jobs(namespace)
	err := jobs.Delete(op.name, deleteSpec)
	if err != nil {
		return fmt.Errorf("Failed to delete job: %s", err)
	}

	return nil
}
