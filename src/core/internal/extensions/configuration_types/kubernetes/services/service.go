package services

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/mujak27/gamen/src/core/internal/extensions/configuration_types/kubernetes/dto"
	"github.com/mujak27/gamen/src/core/internal/models"
	"github.com/mujak27/gamen/src/core/internal/utils"
	uuidUtil "github.com/mujak27/gamen/src/core/utils/uuid"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type KubernetesConfigurationTypeService struct {
	model models.ConfigurationType
}

func NewKubernetesConfigurationTypeService() *KubernetesConfigurationTypeService {

	schema, err := utils.GetSchema(dto.KubernetesConfiguration{})
	if err != nil {
		panic(err)
	}

	id := uuidUtil.GenerateSHA1UUID("configuration-type-kubernetes")

	model := models.ConfigurationType{
		ID:          id,
		Name:        "Kubernetes",
		Description: "Kubernetes is an open-source system for automating deployment, scaling, and management of containerized applications.",
		Schema:      schema,
	}

	return &KubernetesConfigurationTypeService{model: model}
}

func (s *KubernetesConfigurationTypeService) GetId() uuid.UUID {
	return s.model.ID
}

func (s *KubernetesConfigurationTypeService) ValidateConfiguration(configuration models.Configuration) error {

	var config dto.KubernetesConfiguration
	err := utils.ChangeInto(configuration.Data, &config)
	if err != nil {
		return fmt.Errorf("failed to change configuration type: %w", err)
	}
	fmt.Println("config", config)

	validate := validator.New()
	err = validate.Struct(config)
	if err != nil {
		return fmt.Errorf("failed to validate configuration: %w", err)
	}

	return nil
}

// createKubernetesClient creates a Kubernetes client from configuration
func (s *KubernetesConfigurationTypeService) CreateKubernetesClient(config dto.KubernetesConfiguration) (*kubernetes.Clientset, error) {
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
