package services

import (
	"context"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type UtilService struct {
}

func NewUtilService() *UtilService {
	return &UtilService{}
}

// restartDeployment performs the actual deployment restart
func (s *UtilService) RestartDeployment(clientset *kubernetes.Clientset, deploymentName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Get the deployment
	deploy, err := clientset.AppsV1().Deployments("default").Get(ctx, deploymentName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get deployment %s: %w", deploymentName, err)
	}

	// Initialize annotations if nil
	if deploy.Spec.Template.ObjectMeta.Annotations == nil {
		deploy.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
	}

	// Add restart annotation
	deploy.Spec.Template.ObjectMeta.Annotations["kubectl.kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)

	// Update the deployment
	_, err = clientset.AppsV1().Deployments("default").Update(ctx, deploy, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("failed to update deployment %s: %w", deploymentName, err)
	}

	return nil
}
