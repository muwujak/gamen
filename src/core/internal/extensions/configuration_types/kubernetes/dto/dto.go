package dto

// TODO: remove this model and its implementation. change to json based in configuration_type model
type KubernetesConfiguration struct {
	ApiServerEndpoint string `json:"api-server-endpoint" binding:"required"`
	Token             string `json:"token" binding:"required"`
}
