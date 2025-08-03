package interfaces

import "k8s.io/client-go/kubernetes"

type IUtilService interface {
	RestartDeployment(clientset *kubernetes.Clientset, deploymentName string) error
}
