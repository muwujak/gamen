package dto

type KubernetesConfiguration struct {
	ApiServerEndpoint string `json:"api-server-endpoint" binding:"required" validate:"required,url"`
	Token             string `json:"token" binding:"required" validate:"required"`
}
