package dto

// TODO: create base model for all catalogue actions
type RestartDeploymentForm struct {
	DeploymentName  string `json:"deployment_name" binding:"required"`
	ConfigurationId string `json:"configuration_id" binding:"required"`
}
