/* This is an example program to create pods using corev1.Pod */

package main

import (
	"context"
	"fmt"

	testclient "github.com/openshift/cluster-node-tuning-operator/test/e2e/performanceprofile/functests/utils/client"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// main function
func main() {

	// create a varible of type v1.Pod
	var testpod *v1.Pod
	var err error
	var runtimeClass string = "performance-performance"
	//Define Pod
	testpod = &v1.Pod{
		//Define Metadata
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-pod", //Pod name: Explicity specified.
			Namespace: "default",  // The name space where pod should be created
			Labels: map[string]string{
				"name": "test-pod",
			},
			Annotations: map[string]string{
				"irq-load-balancing.crio.io": "housekeeping",
			},
		},
		//Pod specification
		Spec: v1.PodSpec{
			// Containers to run in Pod
			NodeSelector: map[string]string{
				"node-role.kubernetes.io/worker-cnf": "",
			},
			RuntimeClassName: &runtimeClass,
			Containers: []v1.Container{
				{
					Name:    "test-fedora-container",  // container name
					Image:   "fedora:latest",          //container image fedora:latest
					Command: []string{"sleep", "inf"}, // Command that runs which is sleep
					Resources: v1.ResourceRequirements{
						Limits: v1.ResourceList{
							v1.ResourceMemory: resource.MustParse("500M"),
							v1.ResourceCPU:    resource.MustParse("4"),
						},
					},
				},
			},
		},
	}

	err = testclient.Client.Create(context.TODO(), testpod)
	if err != nil {
		fmt.Println("Pod did not get created")
	}
}
