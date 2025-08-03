package interfaces

import (
	"github.com/mujak27/gamen/src/core/internal/extensions/configuration_types/kubernetes/dto"
	"k8s.io/client-go/kubernetes"
)

type IKubernetesService interface {
	CreateKubernetesClient(config dto.KubernetesConfiguration) (*kubernetes.Clientset, error)
}
