package services

import (
	"fmt"

	"github.com/mujak27/gamen/src/core/internal/extensions/configuration_types/kubernetes/dto"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type KubernetesConfigurationService struct {
}

func NewKubernetesConfigurationService() *KubernetesConfigurationService {
	return &KubernetesConfigurationService{}
}

// createKubernetesClient creates a Kubernetes client from configuration
func (s *KubernetesConfigurationService) CreateKubernetesClient(config dto.KubernetesConfiguration) (*kubernetes.Clientset, error) {
	restConfig := &rest.Config{
		Host:            config.ApiServerEndpoint,
		BearerToken:     config.Token,
		TLSClientConfig: rest.TLSClientConfig{Insecure: true},
	}

	fmt.Println("creating kubernetes client for config", config)
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kubernetes client: %w", err)
	}

	return clientset, nil
}
