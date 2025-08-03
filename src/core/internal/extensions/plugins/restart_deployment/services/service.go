package services

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/mujak27/gamen/src/core/internal/api/dto"
	kubernetes_configuration_type "github.com/mujak27/gamen/src/core/internal/extensions/configuration_types/kubernetes/interfaces"
	restart_deployment_interfaces "github.com/mujak27/gamen/src/core/internal/extensions/plugins/restart_deployment/interfaces"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/repository"
	"github.com/mujak27/gamen/src/core/internal/models"
	uuidUtil "github.com/mujak27/gamen/src/core/utils/uuid"

	kubernetes_model "github.com/mujak27/gamen/src/core/internal/extensions/configuration_types/kubernetes/dto"
	"github.com/mujak27/gamen/src/core/internal/extensions/plugins/restart_deployment/utils"
)

// RestartDeploymentPluginService handles deployment restart operations
type RestartDeploymentPluginService struct {
	kubernetesConfigurationService kubernetes_configuration_type.IKubernetesService
	repo                           interfaces.CatalogueRepository
	utilService                    restart_deployment_interfaces.IUtilService
	model                          models.Plugin
}

// NewRestartDeploymentService creates a new instance of RestartDeploymentCatalogueService
// TODO: accept configuration type id as a parameter
func NewRestartDeploymentService(
	kubernetesConfigurationService kubernetes_configuration_type.IKubernetesService,
	repo interfaces.CatalogueRepository,
	utilService restart_deployment_interfaces.IUtilService,
) *RestartDeploymentPluginService {

	id := uuidUtil.GenerateSHA1UUID("plugin-restart-deployment")
	configurationTypeID := uuidUtil.GenerateSHA1UUID("configuration-type-kubernetes")

	model := models.Plugin{
		ID:                  id,
		Name:                "Restart Deployment",
		Type:                models.CatalogueTypeAction,
		UISchema:            models.JSON{},
		Version:             "1.0.0",
		IsActive:            true,
		ConfigurationTypeID: configurationTypeID,
	}

	return &RestartDeploymentPluginService{
		kubernetesConfigurationService: kubernetesConfigurationService,
		repo:                           repo,
		utilService:                    utilService,
		model:                          model,
	}
}

// GetId returns the routing key for restart-deployment service
func (s *RestartDeploymentPluginService) GetId() uuid.UUID {
	return s.model.ID
}

// Get returns the plugin configuration for restart-deployment
func (s *RestartDeploymentPluginService) Get() (models.Plugin, error) {
	return s.model, nil
}

// ValidatePayload validates the incoming payload
func (s *RestartDeploymentPluginService) ActionValidatePayload(payload dto.PluginActionSchema) error {

	_, err := utils.TransformPayloadToModel(payload)
	if err != nil {
		return fmt.Errorf("failed to validate payload: %w", err)
	}

	return nil
}

// Action performs the deployment restart operation
func (s *RestartDeploymentPluginService) Action(payload dto.PluginActionSchema) error {
	log.Println("Starting restart deployment action")

	// Parse and validate pluginData
	pluginData, err := utils.TransformPayloadToModel(payload)
	if err != nil {
		return fmt.Errorf("failed to parse payload: %w", err)
	}

	// Parse configuration data into a model
	var configuration kubernetes_model.KubernetesConfiguration
	if err := utils.Transform(payload.Configuration.Data, &configuration); err != nil {
		return fmt.Errorf("failed to unmarshal configuration fields: %w", err)
	}

	// Create Kubernetes client
	clientset, err := s.kubernetesConfigurationService.CreateKubernetesClient(configuration)
	if err != nil {
		return fmt.Errorf("failed to create Kubernetes client: %w", err)
	}

	// Restart deployment
	if err := s.utilService.RestartDeployment(clientset, pluginData.DeploymentName); err != nil {
		return fmt.Errorf("failed to restart deployment %s: %w", pluginData.DeploymentName, err)
	}

	log.Printf("Successfully restarted deployment: %s", pluginData.DeploymentName)
	return nil
}

// Fetch returns an empty fetch schema as this service doesn't support fetching
func (s *RestartDeploymentPluginService) Fetch(key string) (dto.PluginFetchSchema, error) {
	// TODO: implement fetch logic based on key
	return dto.PluginFetchSchema{}, nil
}
